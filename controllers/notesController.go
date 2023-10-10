package controllers

import (
	"notes-app/models"

	"github.com/gin-gonic/gin"
)

func NotesIndex(context *gin.Context) {
	notes := models.NotesAll()
	context.JSON(200, gin.H{
		"notes": notes,
	})
}
