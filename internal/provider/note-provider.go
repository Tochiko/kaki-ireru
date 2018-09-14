package provider

import (
	"fmt"
	"kaki-ireru/internal/models"
	"strconv"
)


// Find all notes and return an array of notes
func FindNotes(user *models.User) (notes []*models.Note) {
	var test []*models.UserNote
	if err := Db.Where("id_user = ?", strconv.Itoa(user.Id)).Find(&test).Error; err != nil {
		fmt.Println(err)
	} else {
		var keys []int
		for _, userNote := range test {
			keys = append(keys, userNote.IdNote)
		}
		if err := Db.Where(keys).Find(&notes).Error; err != nil {
			fmt.Println(err)
		}
	}
	return
	/*if err := Db.Find(&notes).Error; err != nil {
		fmt.Println(err)
	}
	return*/
}

// Get the specified note or return an error if there is no note
func GetNote(id int) (note models.Note, err error) {
	err = Db.First(&note, id).Error
	return
}

// Create new note and return it
func CreateNote(note *models.Note, user *models.User) *models.Note {
	if err := Db.Create(&note).Error; err != nil {
		fmt.Println(err)
	} else {
		userNote := models.UserNote{user.Id, note.Id}
		if err := Db.Create(&userNote).Error; err != nil {
			fmt.Println(err)
		}
	}



	/*user.Notes = append(user.Notes, note)
	if err := Db.Model(&user).Update("Notes", user.Notes).Error; err != nil {
		fmt.Println(err)
	}*/
	// note.Users = append(note.Users, user)

	/*if err := Db.Create(&note).Error; err != nil {
		fmt.Println(err)
	}*/
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