package controllers

import (
	"github.com/dgrijalva/jwt-go"
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
	claims, _ := c.Get("decoded")
	claimsMap := claims.(jwt.MapClaims)
	idStr, _ := claimsMap["id"].(string)
	id, _ := strconv.Atoi(idStr)
	eMail, _ := claimsMap["eMail"].(string)
	user := models.User{id, eMail, "", ""}

	notes := provider.FindNotes(&user)
	c.JSON(http.StatusOK, notes)
}

// Handler to POST a new note
// The request body is bounded to the note struct and then created
// A new note is the response
func NotesCreationHandler (c *gin.Context) {
	 // var user models.User
	claims, _ := c.Get("decoded")
	claimsMap := claims.(jwt.MapClaims)
	idStr, _ := claimsMap["id"].(string)
	id, _ := strconv.Atoi(idStr)
	eMail, _ := claimsMap["eMail"].(string)
	user := models.User{id, eMail, "", ""}
	var note models.Note
	c.BindJSON(&note)
	provider.CreateNote(&note, &user)
	c.JSON(http.StatusCreated, note)
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
		c.JSON(http.StatusBadRequest, gin.H{"message": "path param id have to be an integer"})
	} else {
		provider.DeleteNote(id)
		c.Status(http.StatusNoContent)
	}


}