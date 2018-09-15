package models

// Note struct is equivalent to the note resource
// Id is an int primary key
type Note struct {
	Id int `gorm:"primary_key;auto_increment"json:"id"`
	Title string `json:"title"binding:"required"`
	Description string `json:"description"binding:"required"`
	Done bool `json:"done"`
}
