package controllers

import (
	"net/http"
	"notes-app/models"
	"notes-app/repositories"
	"notes-app/services"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (u *UserController) Register(c *gin.Context) {
	var input models.UserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	usersRepository := &repositories.UsersRepositoryImpl{}
	signupService := services.SignupService{
		Repository: usersRepository,
	}

	_, err := signupService.Execute(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Registration successful"})
}

func (u *UserController) Login(c *gin.Context) {
	var user models.UserInput
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	usersRepository := &repositories.UsersRepositoryImpl{}
	loginService := services.LoginService{
		Repository: usersRepository,
	}

	token, err := loginService.Execute(user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
