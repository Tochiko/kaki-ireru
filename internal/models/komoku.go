package models

type Komoku struct {
	Id int `gorm:"primary_key;auto_increment"`
	Title string
	Description string
}