package main

import (
	"log"
	"notes-app/controllers"
	"notes-app/models"
	"notes-app/repositories"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.ConnectToDatabase()
	models.MigrateDatabase()

	notesRepository := &repositories.NotesRepositoryImpl{}
	notesController := &controllers.NotesController{
		Repository: notesRepository,
	}

	r.GET("/notes", notesController.Index)
	r.POST("/notes", notesController.Create)

	log.Println("Starting server on port 8080")

	r.Run()
}
