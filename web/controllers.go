package web

import (
	"github.com/gin-gonic/gin"
	"kaki-ireru/internal/provider"
	"net/http"
)

func notesListHandler (c *gin.Context) {
	notes := provider.FindNotes()
	c.JSON(http.StatusOK, notes)
}