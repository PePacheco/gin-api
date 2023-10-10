package services

import (
	"errors"
	"notes-app/helpers"
	"notes-app/models"
	"notes-app/repositories"
)

type SignupService struct {
	Repository repositories.UsersRepository
}

func (self *SignupService) Execute(userInput models.UserInput) (*models.User, error) {
	hashedPassword, err := helpers.HashPassword(userInput.Password)
	if err != nil {
		return nil, errors.New("Error hashing password")
	}

	user := models.User{Username: userInput.Username, Email: userInput.Email, Password: hashedPassword}

	err = self.Repository.Create(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
