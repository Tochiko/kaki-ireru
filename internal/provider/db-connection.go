package provider

import (
	"github.com/jinzhu/gorm"
	"kaki-ireru/internal/models"
)

var Db *gorm.DB


// Initialize the database var from provider
func InitDatabase (connectionPool *gorm.DB) {
	Db = connectionPool
	Db.AutoMigrate(models.Note{})
	Db.AutoMigrate(models.User{})
	Db.AutoMigrate(models.UserNote{})
}