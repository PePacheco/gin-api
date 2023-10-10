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

	r.GET("/notes", controllers.NotesIndex)

	log.Println("Starting server on port 8080")

	r.Run()
}
