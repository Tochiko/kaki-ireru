package web

import "github.com/gin-gonic/gin"

// Registry routes and http methods to given gin router
// Corresponding handler functions are referred
func RegistryRoutes (router gin.IRouter) {
	router.GET("/notes", notesFindHandler)
	router.POST("/notes", notesCreationHandler)
	router.PUT("/notes", notesReplaceHandler)

	router.GET("/notes/:id", notesFindByIdHandler)
	router.DELETE("/notes/:id", notesDeletionHandler)
}