package controllers

import (
	"notes-app/models"
	"notes-app/repositories"

	"github.com/gin-gonic/gin"
)

type NotesController struct {
	Repository repositories.NotesRepository
}

func (c *NotesController) Index(context *gin.Context) {
	notes := c.Repository.All()
	context.JSON(200, gin.H{
		"notes": notes,
	})
}

func (c *NotesController) Create(context *gin.Context) {
	var input models.NoteInput
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	note := models.Note{Name: input.Name, Content: input.Content}
	c.Repository.Create(&note)
	context.JSON(200, gin.H{
		"note": note,
	})
}
