package model

import (
	"time"
)

// User represents a user model.
type User struct {
	ID          uint `gorm:"primaryKey"`
	Name        string
	Email       string `gorm:"uniqueIndex"`
	DateOfBirth time.Time
}
