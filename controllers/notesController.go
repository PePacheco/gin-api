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

func NotesCreate(context *gin.Context) {
	var input models.NoteInput
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	note := models.Note{Name: input.Name, Content: input.Content}
	models.NotesCreate(&note)
	context.JSON(200, gin.H{
		"note": note,
	})
}
