package repository

import (
	"github.com/dolearci/go-microservice/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

// FindAll finds all users.
func (r *UserRepository) FindAll() ([]model.User, error) {
	var users []model.User
	result := r.db.Find(&users)
	return users, result.Error
}
