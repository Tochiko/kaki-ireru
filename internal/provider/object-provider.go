package provider

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"kaki-ireru/internal/models"
)

var db *gorm.DB

func InitDatabase (connectionPool *gorm.DB) {
	db = connectionPool
	db.AutoMigrate(models.Komoku{})
}

func FindKomokus () (komokus []*models.Komoku) {
	if err := db.Find(&komokus).Error; err != nil {
		fmt.Println(err)
	}
	return
}

func GetKomoku (id int) (komoku models.Komoku) {
	if err := db.First(&komoku, id).Error; err != nil {
		fmt.Println(err)
	}
	return
}

func CreateKomoku (komoku *models.Komoku) *models.Komoku {
	if err := db.Create(&komoku).Error; err != nil {
		fmt.Println(err)
	}
	return komoku
}

func UpdateKomoku (komoku *models.Komoku) *models.Komoku {
	if err := db.Create(&komoku).Error; err != nil {
		fmt.Println(err)
	}
	return komoku
}

func DeleteKomoku (id int) {
	if err := db.Delete(&models.Komoku{id, nil, nil}).Error; err != nil {
		fmt.Println(err)
	}
}