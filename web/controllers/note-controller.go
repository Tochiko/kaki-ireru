package controllers

import (
	"github.com/gin-gonic/gin"
	"kaki-ireru/internal/models"
	"kaki-ireru/internal/provider"
	"net/http"
	"strconv"
)

// Handler to GET a list resource of .../notes
// An array of note objects will be respond
// If not notes are found the response contains an empty array
func NotesFindHandler(c *gin.Context) {
	user, exists := getUserFromContext(c)
	if !exists {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "something's gone wrong please try it again"})
		return
	}
	notes := provider.FindNotes(user)
	c.JSON(http.StatusOK, notes)
}

// Handler to POST a new note
// The request body is bounded to the note struct and then created
// A new note is the response
func NotesCreationHandler (c *gin.Context) {
	user, exists := getUserFromContext(c)
	if !exists {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "something's gone wrong please try it again"})
		return
	}
	var note models.Note
	c.BindJSON(&note)
	provider.CreateNote(&note, user)
	c.JSON(http.StatusCreated, &note)
	return
}

// Handler to PUT a note
// The request body is bounded tho the note struct and then replaced
// Replaced note is the response
func NotesReplaceHandler (c *gin.Context) {
	var note models.Note
	c.BindJSON(&note)
	provider.UpdateNote(&note)
	c.JSON(http.StatusOK, note)
}

// Get a specified note by id
// If the id isn't an int a 400 is respond
// If the id does not refer a note a 404 is respond
// Else the response contains the requested note
func NotesFindByIdHandler (c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "path param id have to be an integer"})
	} else {
		note, e := provider.GetNote(id)
		if e != nil {
			c.JSON(http.StatusNotFound, gin.H{"message": e.Error()})
		} else {
			c.JSON(http.StatusOK, note)
		}
	}
}

// Delete a note by their id
// If the id isn't an int a 400 is respond
// Else the node will be deleted
func NotesDeletionHandler (c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "path param id have to be an integer"})
		return
	}
	user, exists := getUserFromContext(c)
	if !exists {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "something's gone wrong please try it again"})
		return
	}

	note := &models.Note{Id: id, Title: "", Description: "", Done: false}
	provider.DeleteNote(note, user)
	c.Status(http.StatusNoContent)
	return
}

func getUserFromContext (c *gin.Context) (*models.User, bool){
	u, exists := c.Get("user")
	user := u.(*models.User)
	return user, exists
}