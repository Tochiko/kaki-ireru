package models

type Note struct {
	Id int `gorm:"primary_key;auto_increment"json:"id"binding:"required"`
	Title string `json:"title"binding:"required"`
	Description string `json:"description"binding:"required"`
}