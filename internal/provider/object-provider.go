package provider

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"kaki-ireru/internal/models"
)

var db *gorm.DB

// Initialize the database var from provider
func InitDatabase (connectionPool *gorm.DB) {
	db = connectionPool
	db.AutoMigrate(models.Note{})
}

// Find all notes and return an array of notes
func FindNotes() (notes []*models.Note) {
	if err := db.Find(&notes).Error; err != nil {
		fmt.Println(err)
	}
	return
}

// Get the specified note or return an error if there is no note
func GetNote(id int) (note models.Note, err error) {
	err = db.First(&note, id).Error
	return
}

// Create new note and return it
func CreateNote(note *models.Note) *models.Note {
	if err := db.Create(&note).Error; err != nil {
		fmt.Println(err)
	}
	return note
}

// Update existing note
func UpdateNote(note *models.Note) *models.Note {
	if err := db.Save(&note).Error; err != nil {
		fmt.Println(err)
	}
	return note
}

// Delete a note by given id
func DeleteNote(id int) {
	if err := db.Delete(&models.Note{id, "", "", false}).Error; err != nil {
		fmt.Println(err)
	}
}