package main

import (
	"log"
	"notes-app/controllers"
	"notes-app/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.ConnectToDatabase()
	models.MigrateDatabase()
	notesController := &controllers.NotesController{}

	r.GET("/notes", notesController.Index)
	r.POST("/notes", notesController.Create)

	log.Println("Starting server on port 8080")

	r.Run()
}
