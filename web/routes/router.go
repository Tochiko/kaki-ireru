package routes

import (
	"github.com/gin-gonic/gin"
	"kaki-ireru/web/controllers"
	"kaki-ireru/web/middleware"
)

func RegistryRoutes (router gin.IRouter) {
	notes := router.Group("/notes")
	notes.Use(middleware.TokenDecoding)

	notes.GET("/", controllers.NotesFindHandler)
	notes.POST("/", controllers.NotesCreationHandler)
	notes.PUT("/", controllers.NotesReplaceHandler)

	notes.GET("/:id", controllers.NotesFindByIdHandler)
	notes.DELETE("/:id", controllers.NotesDeletionHandler)
}