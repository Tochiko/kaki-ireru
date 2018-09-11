package controllers

import (
	"github.com/gin-gonic/gin"
	"kaki-ireru/internal/models"
	"kaki-ireru/internal/provider"
	"net/http"
)

func CreateUser (c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	err := provider.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	} else {
		// todo: test the different options
		c.Writer.WriteHeader(http.StatusNoContent)
		// c.AbortWithStatus(http.StatusCreated)
		// c.Status(http.StatusNoContent)
	}
}

func VerifyUser (c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	verification := provider.VerifyUser(&user)
	if verification {
		c.Status(http.StatusOK)
	} else {
		c.Status(http.StatusUnauthorized)
	}
}