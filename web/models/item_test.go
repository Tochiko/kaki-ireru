package models_test

import (
	"errors"
	"kaki-ireru/web/models"
)


var store = map[string]*models.User{
	"U1":
		{Id: "U1", Items: []*models.Item{
		{Id: 1, Title: "First", Description: "Item for some tests", Done: false},
		{Id: 2, Title: "Second", Description: "Item for some tests", Done: true},
		{Id: 3, Title: "Third", Description: "Item for some tests", Done: false},
	}},
	"U2":
		{Id: "U2", Items: []*models.Item{
		{Id: 4, Title: "Fourth", Description: "Item for some tests", Done: true},
		{Id: 5, Title: "Fifth", Description: "Item for some tests", Done: true},
		{Id: 6, Title: "Sixth", Description: "Item for some tests", Done: false},
	}},
	"U3":
		{Id: "U3", Items: []*models.Item{
		{Id: 7, Title: "Seventh", Description: "Item for some tests", Done: false},
		{Id: 8, Title: "Eighth", Description: "Item for some tests", Done: false},
		{Id: 9, Title: "Ninth", Description: "Item for some tests", Done: false},
	}},
}

func (db *TestDb) AllItems (user *models.User) ([]*models.Item, error) {
	var items []*models.Item
	data := store[user.Id]
	if data != nil {
		items = data.Items
	}
	return items, nil
}

func (db *TestDb) ItemById (user *models.User, id int) (*models.Item, error) {
	data := store[user.Id]
	err := errors.New("no item found")

	if data == nil {
		return nil, err
	}

	for _, item := range data.Items {
		if item.Id == id {
			return item, nil
		}
	}
	return nil, err
}

func (db *TestDb) CreateItem (user *models.User, item *models.Item) (*models.Item, error) {
	exist := store[user.Id] != nil

	if !exist {
		user.Items = append(user.Items, item)
		store[user.Id] = user
	}

	store[user.Id].Items = append(store[user.Id].Items, item)
	return item, nil
}