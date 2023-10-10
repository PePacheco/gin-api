package models

import (
	"log"
	"time"
)

type User struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Username  string `gorm:"size:255;uniqueIndex;not null"`
	Password  string `gorm:"size:255;not null"`
	Email     string `gorm:"size:255;uniqueIndex;not null"`
}

type UserInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

func UserByUsername(username string) *User {
	var user User
	result := database.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return nil
	}
	return &user
}

func UserCreate(user *User) error {
	result := database.Create(user)
	if result.Error != nil {
		log.Println("Database error:", result.Error) // Print more detailed error
	}
	return result.Error
}
