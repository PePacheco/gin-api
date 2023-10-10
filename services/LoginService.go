package services

import (
	"errors"
	"notes-app/helpers"
	"notes-app/models"
	"notes-app/repositories"
)

type LoginService struct {
	Repository repositories.UsersRepository
}

func (self *LoginService) Execute(userInput models.UserInput) (string, error) {
	foundUser := self.Repository.ShowByUsername(userInput.Username)
	if foundUser == nil {
		return "", errors.New("User not found")
	}

	if !helpers.CheckPassword(userInput.Password, foundUser.Password) {
		return "", errors.New("Incorrect password")
	}

	token, err := helpers.CreateToken(foundUser)
	if err != nil {
		return "", err
	}

	return token, nil
}
