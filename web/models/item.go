package models

import "errors"

type Item struct {
	Id int `gorm:"primary_key;auto_increment"json:"id"`
	Title string `json:"title"binding:"required"`
	Description string `json:"description"binding:"required"`
	Done bool `json:"done"`
}

func (db *DB) AllItems (u *User) (items []*Item, err error) {
	err = db.Model(&u).Association("Items").Find(&items).Error
	return
}

func (db *DB) ItemById (u *User, id int) (*Item, error) {
	err := db.Preload("Items", "id = ?", id).Find(&u).Error
	if err != nil {
		return nil, err
	}

	if len(u.Items) <= 0 {
		err = errors.New("no items found for given id")
		return nil, err
	}

	return u.Items[0], nil
}

func (db *DB) CreateItem (u *User, item *Item) (*Item, error) {
	err := db.Model(&u).Association("Items").Append(item).Error
	return item, err
}