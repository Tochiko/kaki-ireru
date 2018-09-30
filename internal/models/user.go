package models

type User struct {
	Id string `gorm:"primary_key"json:"id"`

	// Many to Many relation trough user_notes linkage
	Notes []*Note `gorm:"many2many:user_notes;"`
}