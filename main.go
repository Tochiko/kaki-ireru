package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"kaki-ireru/internal/provider"
	"kaki-ireru/web/routes"
	"os"
)

// First: Open db connection pool
// Second: Initialize the provider from internal package
// Third: Create the gin router and run it
func main () {
	db, err := gorm.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		panic("failed to connect to database")
	}
	defer db.Close()
	// todo: remove the following line - is only for some tests
	/*db.Exec("DROP TABLE users;")
	db.Exec("DROP TABLE notes;")
	db.Exec("DROP TABLE user_notes;")*/
	provider.InitDatabase(db)

	router := gin.Default()
	routes.RegistryRoutes(router)
	router.Run()
}
