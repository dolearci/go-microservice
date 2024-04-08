package main

import (
	"github.com/dolearci/go-microservice/config"
	handlers "github.com/dolearci/go-microservice/handler"
	"github.com/dolearci/go-microservice/repository"
	"github.com/gin-gonic/gin"
)

func main() {
	database := config.ConnectToDatabase()

	router := gin.Default()

	userRepository := repository.NewUserRepository(database)
	userHandler := handlers.NewUserHandler(userRepository)

	router.GET("/", userHandler.GetAllUsers)
	router.GET("/:id", userHandler.GetUserByID)
	router.POST("/save", userHandler.CreateUser)

	if err := router.Run(":8080"); err != nil {
		panic("Failed to start the server")
	}
}
