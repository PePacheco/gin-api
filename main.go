package main

import (
	"log"
	"notes-app/controllers"
	"notes-app/middlewares"
	"notes-app/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.ConnectToDatabase()
	models.MigrateDatabase()

	notesController := &controllers.NotesController{}
	r.GET("/notes", middlewares.AuthMiddleware(), notesController.Index)
	r.POST("/notes", middlewares.AuthMiddleware(), notesController.Create)
	r.GET("/notes/:id", middlewares.AuthMiddleware(), notesController.Show)
	r.DELETE("/notes/:id", middlewares.AuthMiddleware(), notesController.Delete)

	usersController := &controllers.UserController{}
	r.POST("/register", usersController.Register)
	r.POST("/login", usersController.Login)

	log.Println("Starting server on port 8080")

	r.Run()
}
