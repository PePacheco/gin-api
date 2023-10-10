package controllers

import (
	"notes-app/models"
	"notes-app/repositories"
	"notes-app/services"

	"github.com/gin-gonic/gin"
)

type NotesController struct{}

func (c *NotesController) Index(context *gin.Context) {
	notesRepository := &repositories.NotesRepositoryImpl{}
	service := services.ListAllNotesService{
		Repository: notesRepository,
	}
	notes := service.Execute()
	context.JSON(200, gin.H{
		"notes": notes,
	})
}

func (c *NotesController) Create(context *gin.Context) {
	notesRepository := &repositories.NotesRepositoryImpl{}
	var input models.NoteInput
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	note := models.Note{Name: input.Name, Content: input.Content}
	service := services.CreateNoteService{
		Repository: notesRepository,
	}
	service.Execute(&note)
	context.JSON(200, gin.H{
		"note": note,
	})
}
