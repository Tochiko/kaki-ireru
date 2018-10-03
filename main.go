package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"kaki-ireru/web/controllers"
	"kaki-ireru/web/middleware"
	"kaki-ireru/web/models"
	"os"
)

func main () {
	db := models.NewDB(os.Getenv("DATABASE_URL"))
	defer db.Close()

	env := controllers.Env{Db: db}
	router := gin.Default()

	notes := router.Group("/items")
	notes.Use(middleware.TokenDecoding)
	notes.GET("/", env.AllItems)
	notes.GET("/:id", env.ItemById)
	notes.POST("/", env.CreateItem)

	/*db.Exec("DROP TABLE users;")
	db.Exec("DROP TABLE notes;")
	db.Exec("DROP TABLE user_notes;")*/

	router.Run()
}
