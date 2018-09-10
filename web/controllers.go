package web

import (
	"github.com/gin-gonic/gin"
	"kaki-ireru/internal/models"
	"kaki-ireru/internal/provider"
	"net/http"
	"strconv"
)

func notesFindHandler(c *gin.Context) {
	notes := provider.FindNotes()
	c.JSON(http.StatusOK, notes)
}

func notesCreationHandler (c *gin.Context) {
	var note models.Note
	c.BindJSON(&note)
	provider.CreateNote(&note)
	c.JSON(http.StatusCreated, note)
}

func notesReplaceHandler (c *gin.Context) {
	var note models.Note
	c.BindJSON(&note)
	provider.UpdateNote(&note)
	c.JSON(http.StatusOK, note)
}

func notesFindByIdHandler (c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "path param id have to be an integer"})
	} else {
		note := provider.GetNote(id) // NOT FOUND isn't handled at this time!
		c.JSON(http.StatusOK, note)
	}
}

func notesDeletionHandler (c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "path param id have to be an integer"})
	} else {
		provider.DeleteNote(id)
		c.Status(http.StatusNoContent)
	}


}