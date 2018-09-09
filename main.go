package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"kaki-ireru/internal/provider"
	"kaki-ireru/web"
	"os"
)

var db *gorm.DB
var err error

func main () {
	db, _ = gorm.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		panic("failed to connect to database")
	}
	defer db.Close()
	provider.InitDatabase(db)

	router := gin.Default()
	web.RegistryRoutes(router)
	router.Run()
}
