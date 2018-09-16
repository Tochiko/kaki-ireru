package provider

import (
	"kaki-ireru/internal/models"
)


// Find all notes and return an array of notes
func FindNotes(user *models.User) (notes []*models.Note, err error) {
	err = Db.Model(&user).Related(&notes, "Notes").Error
	return
}

// Get the specified note or return an error if there is no note
func GetNote(id int) (note models.Note, err error) {
	err = Db.First(&note, id).Error
	return
}

// Create new note and return it
func CreateNote(note *models.Note, user *models.User) (*models.Note, error) {
	// first create the note
	if err := Db.Create(note).Error; err != nil {
		return nil, err
	}
	// second relate the created note with user
	if err := Db.Model(&user).Association("Notes").Append(note).Error; err != nil {
		return nil, err // TODO normally the already created note can be deleted then
	}
	// if everything worked well then return the note and nil
	return note, nil
}

// Update existing note
func UpdateNote(note *models.Note) (*models.Note, error) {
	if err := Db.Save(&note).Error; err != nil {
		return nil, err
	}
	return note, nil
}

// Delete a note by given id
func DeleteNote(note *models.Note, user *models.User) (err error) {
	// first: delete the relation between note and user
	err = Db.Model(&user).Association("Notes").Delete(note).Error
	if err != nil {
		return
	}
	// second: delete the note - if everything worked well the returned error is nil
	err = Db.Delete(note).Error
	return
}