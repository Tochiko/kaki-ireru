package provider

import (
	"fmt"
	"kaki-ireru/internal/models"
)


// Find all notes and return an array of notes
func FindNotes() (notes []*models.Note) {
	if err := Db.Find(&notes).Error; err != nil {
		fmt.Println(err)
	}
	return
}

// Get the specified note or return an error if there is no note
func GetNote(id int) (note models.Note, err error) {
	err = Db.First(&note, id).Error
	return
}

// Create new note and return it
func CreateNote(note *models.Note) *models.Note {
	if err := Db.Create(&note).Error; err != nil {
		fmt.Println(err)
	}
	return note
}

// Update existing note
func UpdateNote(note *models.Note) *models.Note {
	if err := Db.Save(&note).Error; err != nil {
		fmt.Println(err)
	}
	return note
}

// Delete a note by given id
func DeleteNote(id int) {
	if err := Db.Delete(&models.Note{id, "", "", false}).Error; err != nil {
		fmt.Println(err)
	}
}