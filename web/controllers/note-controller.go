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
	// find the referenced notes for a user
	user := getUserFromClaims(c)
	notes, err := provider.FindNotes(user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	// if everything worked well then respond an array of notes with status ok
	c.JSON(http.StatusOK, notes)
}

// Handler to POST a new note
// The request body is bounded to the note struct and then created
// A new note is the response
func NotesCreationHandler (c *gin.Context) {
	// get the corresponding user object from user fields in the claims
	user := getUserFromClaims(c)
	// bind the json input into a new note object
	var note models.Note
	c.BindJSON(&note)
	// create the note object and refer it to corresponding user
	_, err := provider.CreateNote(&note, user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	// if everything worked well then respond the created note with status created
	c.JSON(http.StatusCreated, &note)
}

// Handler to PUT a note
// The request body is bounded tho the note struct and then replaced
// Replaced note is the response
func NotesReplaceHandler (c *gin.Context) { // TODO consider user by replacing notes
	// bind the json input into a new note object
	var note models.Note
	c.BindJSON(&note)
	_, err := provider.UpdateNote(&note)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	// if everything worked well then respond the new note with status ok
	c.JSON(http.StatusOK, note)
}

// Get a specified note by id
// If the id isn't an int a 400 is respond
// If the id does not refer a note a 404 is respond
// Else the response contains the requested note
func NotesFindByIdHandler (c *gin.Context) { // TODO consider user by get a note by id
	// extract the id from path and convert it into an integer
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "the path param id must be an integer"})
		return
	}
	// get the note from provider - if there's an error then no note has been found
	note, err := provider.GetNote(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	// if everything worked well then respond the note with status ok
	c.JSON(http.StatusOK, note)
}

// Delete a note by their id
// If the id isn't an int a 400 is respond
// Else the node will be deleted
func NotesDeletionHandler (c *gin.Context) {
	// extract the id from path and convert it into an integer
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "the path param id must be an integer"})
		return
	}
	// create user object corresponding to the claims from context and a note corresponding to the id from path
	user := getUserFromClaims(c)
	note := &models.Note{Id: id, Title: "", Description: "", Done: false}
	// delete the relation and the note
	err = provider.DeleteNote(note, user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	// if everything worked well then respond with status no content
	c.Status(http.StatusNoContent)
}

/**
extract user fields from decoded claims in context
return a created user object based on extracted fields
 */
func getUserFromClaims (c *gin.Context) (user *models.User) {
	// check if claims exist in decoded key from context
	claims, exists := c.Get("decoded")
	if !exists {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "the received token does not contain claims"})
		return
	}
	// parse the claims into type jwt.MapClaims
	claimsMap := claims.(jwt.MapClaims)
	// get the value from the id field from claims
	// mapClaims are typed as map[string]interface{} and the id is parsed from json so it's typed as float64
	// to get the id as int use type assertion to receive a float64 from the claimsMap and create an int with it
	id := int(claimsMap["id"].(float64))
	// create a user object with id and set it to context
	user = &models.User{id, "", "", "", nil}
	return
}