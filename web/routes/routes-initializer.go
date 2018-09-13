package routes

import "github.com/gin-gonic/gin"

func RegistryRoutes (router gin.IRouter) {
	registryNoteRoutes(router)
	registryUserRoutes(router)
}