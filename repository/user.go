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

func (repository *UserRepository) FindAll() ([]model.User, error) {
	var users []model.User
	result := repository.db.Find(&users)
	return users, result.Error
}

func (repository *UserRepository) Create(user *model.User) error {
	return repository.db.Create(&user).Error
}

func (repository *UserRepository) FindByID(id uint) (model.User, error) {
	var user model.User
	result := repository.db.First(&user, id)
	return user, result.Error
}

func (repository *UserRepository) DeleteByID(id uint) error {
	return repository.db.Delete(&model.User{}, id).Error
}

func (repository *UserRepository) DeleteAll() error {
	return repository.db.Where("id > ?", 0).Delete(&model.User{}).Error
}

func (repository *UserRepository) Update(id uint, user model.User) error {
	var existingUser model.User
	if err := repository.db.First(&existingUser, id).Error; err != nil {
		return err
	}

	return repository.db.Model(&existingUser).Updates(user).Error
}
