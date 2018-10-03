package models

type User struct {
	Id string `gorm:"primary_key"json:"id"`
	Items []*Item `gorm:"many2many:user_items;"`
}
