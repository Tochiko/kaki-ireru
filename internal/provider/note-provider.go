package provider

import (
	"fmt"
	"kaki-ireru/internal/models"
)


// Find all notes and return an array of notes
func FindNotes(user *models.User) (notes []*models.Note) {
	if err := Db.Model(&user).Related(&notes, "Notes").Error; err != nil {
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
func CreateNote(note *models.Note, user *models.User) *models.Note {
	if err := Db.Create(note).Error; err != nil {
		fmt.Println(err)
	} else {
		if err := Db.Model(&user).Association("Notes").Append(note).Error; err != nil {
			fmt.Println(err)
		}
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
func DeleteNote(note *models.Note, user *models.User) {
	if err := Db.Model(&user).Association("Notes").Delete(note).Error; err != nil {
		fmt.Println(err)
	} else {
		if err := Db.Delete(note).Error; err != nil {
			fmt.Println(err)
		}
	}
}