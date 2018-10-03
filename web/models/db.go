package models

import (
	"github.com/jinzhu/gorm"
)

type DataStore interface {
	AllItems (user *User) ([]*Item, error)
	ItemById (user *User, id int) (*Item, error)
	CreateItem (user *User, item *Item) (*Item, error)
}

type DB struct {
	*gorm.DB
}

func NewDB (url string) (*DB) {
	db, err := gorm.Open("postgres", url)
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Item{})
	db.AutoMigrate(&User{})

	return &DB{DB: db}
}