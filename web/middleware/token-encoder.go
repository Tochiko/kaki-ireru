package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strings"
)

// Middleware function for gin gonic
// Bearer tokens are decoded and set to context
// In cases of errors or invalid tokens the request cycle is interrupted with corresponding status and message
func TokenDecoding (c *gin.Context) {
	c.Header("WWW-Authenticate", "Basic http://localhost:3000/login")

	authString := c.GetHeader("Authorization")
	if authString == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "you have to be authorized to use this service"})
		return
	}

	authValues := strings.SplitN(authString, " ", 2)
	if len(authValues) != 2 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "the authorization values are wrong structured"})
		return
	}

	token, err := jwt.Parse(authValues[1], func (token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("token has wrong signing method")
		}
		return []byte(os.Getenv("HS_PRIVATE_KEY")), nil
	})

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "this token is not allowed"})
		return
	}

	if !token.Valid {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "the token is invalid"})
		return
	} else {
		c.Header("WWW-Authenticate", "")
		c.Set("decoded", token.Claims)
	}
}
