package models

type Note struct {
	Id int `gorm:"primary_key;auto_increment"`
	Title string
	Description string
}