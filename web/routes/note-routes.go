package routes

import (
	"github.com/gin-gonic/gin"
	"kaki-ireru/web/controllers"
)

// Registry routes and http methods to given gin router
// Corresponding handler functions are referred
func registryNoteRoutes(router gin.IRouter) {
	router.GET("/notes", controllers.NotesFindHandler)
	router.POST("/notes", controllers.NotesCreationHandler)
	router.PUT("/notes", controllers.NotesReplaceHandler)

	router.GET("/notes/:id", controllers.NotesFindByIdHandler)
	router.DELETE("/notes/:id", controllers.NotesDeletionHandler)
}