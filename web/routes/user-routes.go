package routes

import (
	"github.com/gin-gonic/gin"
	"kaki-ireru/web/controllers"
)

func registryUserRoutes (router gin.IRouter) {
	router.POST("/users", controllers.CreateUser)
	router.POST("/login", controllers.VerifyUser)
}
