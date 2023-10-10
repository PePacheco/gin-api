package models

import (
	"time"

	"gorm.io/gorm"
)

type Note struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"size:255;not null"`
	Content   string `gorm:"type:text"`
	CreatedAt time.Time
	UpdatedAt time.Time      `gorm:"index"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type NoteInput struct {
	Name    string `json:"name" binding:"required"`
	Content string `json:"content" binding:"required"`
}

func NotesAll() *[]Note {
	var notes []Note
	database.Where("deleted_at is NULL").Order("updated_at desc").Find(&notes)
	return &notes
}

func NotesCreate(note *Note) {
	database.Create(note)
}

func NoteById(id uint) *Note {
	var note Note
	result := database.First(&note, id)
	if result.Error != nil {
		return nil
	}
	println(note.Name)
	return &note
}

func NoteByName(name string) *Note {
	var note Note
	result := database.Where("name = ?", name).First(&note)
	if result.Error != nil {
		return nil
	}
	return &note
}

func NotesDelete(note *Note) {
	database.Delete(note)
}
