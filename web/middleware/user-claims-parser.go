package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"kaki-ireru/internal/models"
	"strconv"
)

func ParseUserClaims (c *gin.Context) {
	claims, exists := c.Get("decoded")
	if !exists {
		return
	}
	claimsMap := claims.(jwt.MapClaims)
	idStr, _ := claimsMap["id"].(string)
	id, _ := strconv.Atoi(idStr)
	eMail, _ := claimsMap["eMail"].(string)
	user := models.User{id, eMail, "", "", nil}
	c.Set("user", &user)
}