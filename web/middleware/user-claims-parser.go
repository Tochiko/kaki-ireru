package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"kaki-ireru/internal/models"
	"net/http"
)

func ParseUserClaims (c *gin.Context) {
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
	user := models.User{id, "", "", "", nil}
	c.Set("user", &user)
}