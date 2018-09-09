package web

import "github.com/gin-gonic/gin"

func RegistryRoutes (router gin.IRouter) {
	router.GET("/notes", notesListHandler)
}