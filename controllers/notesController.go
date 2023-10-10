package controllers

import (
	"notes-app/models"
	"notes-app/repositories"
	"notes-app/services"
	"strconv"

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
	err := service.Execute(&note)
	if err != nil {
		context.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	context.JSON(200, gin.H{
		"note": note,
	})
}

func (c *NotesController) Show(context *gin.Context) {
	notesRepository := &repositories.NotesRepositoryImpl{}
	id, _ := strconv.ParseUint(context.Param("id"), 10, 64)
	service := services.ShowNoteService{
		Repository: notesRepository,
	}
	note := service.Execute(uint(id))
	context.JSON(200, gin.H{
		"note": note,
	})
}
