package provider

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"kaki-ireru/internal/models"
)

var db *gorm.DB

func InitDatabase (connectionPool *gorm.DB) {
	db = connectionPool
	db.AutoMigrate(models.Note{})
}

func FindNotes() (notes []*models.Note) {
	if err := db.Find(&notes).Error; err != nil {
		fmt.Println(err)
	}
	return
}

func GetNote(id int) (note models.Note) {
	if err := db.First(&note, id).Error; err != nil {
		fmt.Println(err)
	}
	return
}

func CreateNote(note *models.Note) *models.Note {
	if err := db.Create(&note).Error; err != nil {
		fmt.Println(err)
	}
	return note
}

func UpdateNote(note *models.Note) *models.Note {
	if err := db.Save(&note).Error; err != nil {
		fmt.Println(err)
	}
	return note
}

func DeleteNote(id int) {
	if err := db.Delete(&models.Note{id, "", ""}).Error; err != nil {
		fmt.Println(err)
	}
}