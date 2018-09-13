package routes

import (
	"github.com/gin-gonic/gin"
	"kaki-ireru/web/controllers"
	"kaki-ireru/web/middleware"
)

// Registry routes and http methods to given gin router
// Corresponding handler functions are referred
func registryNoteRoutes(router gin.IRouter) {
	notes := router.Group("/notes")
	notes.Use(middleware.TokenDecoding)

	notes.GET("/", controllers.NotesFindHandler)
	notes.POST("/", controllers.NotesCreationHandler)
	notes.PUT("/", controllers.NotesReplaceHandler)

	notes.GET("/:id", controllers.NotesFindByIdHandler)
	notes.DELETE("/:id", controllers.NotesDeletionHandler)
}