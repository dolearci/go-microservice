// integration_test.go
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/dolearci/go-microservice/config"
	"github.com/dolearci/go-microservice/model"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestCreateAndGetUser(t *testing.T) {
	router := config.SetupRouter()
	server := httptest.NewServer(router)
	defer server.Close()

	testEmail := "test" + fmt.Sprintf("%d", time.Now().Unix()) + "@example.com"

	user := model.User{
		Name:        "Test User",
		Email:       testEmail,
		DateOfBirth: time.Now(),
	}

	userJSON, _ := json.Marshal(user)
	resp, err := http.Post(fmt.Sprintf("%s/save", server.URL), "application/json", bytes.NewBuffer(userJSON))
	if err != nil || resp.StatusCode != http.StatusOK {
		t.Fatalf("Failed to create user: %v", err)
	}

	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/1", server.URL), nil)
	resp, err = http.DefaultClient.Do(req)
	if err != nil || resp.StatusCode != http.StatusOK {
		t.Fatalf("Failed to get user: %v", err)
	}

	var fetchedUser model.User
	if err := json.NewDecoder(resp.Body).Decode(&fetchedUser); err != nil || fetchedUser.ID != 1 {
		t.Fatalf("User fetched does not match expected ID")
	}
}
