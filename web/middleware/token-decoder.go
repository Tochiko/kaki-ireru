package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/kataras/iris/core/errors"
	"net/http"
	"os"
	"strings"
)

// Middleware function for gin gonic
// Bearer tokens are decoded and set to context
// In cases of errors or invalid tokens the request cycle is interrupted with corresponding status and message
func TokenDecoding (c *gin.Context) {
	c.Header("WWW-Authenticate", "Basic http://localhost:3000/login")
	// extract the token as string from authorization header
	tokenString, err := extractTokenString(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": err.Error()})
		return
	}
	// parse the tokenString to *jwt.Token
	token, err := parseToken(tokenString)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": err.Error()})
		return
	}
	// if everything worked well remove the www-authenticate header and set token claims to context
	c.Header("WWW-Authenticate", "")
	c.Set("decoded", token.Claims)
}

/**
extract token strings from http authorization header
in case of errors a descriptive message is returned
 */
func extractTokenString (c *gin.Context) (tokenString string, err error) {
	errorMessage := "you have to set a valid bearer token to use this service"
	// get the authorization header from context - if it's empty an error is returned
	authString := c.GetHeader("Authorization")
	if authString == "" {
		err = errors.New(errorMessage)
		return
	}
	// split the header in two parts: [bearer, tokenString]
	authValues := strings.SplitN(authString, " ", 2)
	if len(authValues) != 2 {
		err = errors.New(errorMessage)
		return
	}
	// the index in which is the token as string
	tokenString = authValues[1]
	return
}

/**
parse tokenString into *jwt.Token
in case of errors a descriptive message is returned
 */
func parseToken (tokenString string) (token *jwt.Token, err error) {
	errorMessage := "the received token is not allowed"
	// parse the token string into a jwt token
	// the inner function is initially checking if tokens signing method can be cast with SigningMethodHMAC
	// if the cast has success so tokens method is ok and the private key is returned from inner function
	token, err = jwt.Parse(tokenString, func (t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf(errorMessage)
		}
		return []byte(os.Getenv("HS_PRIVATE_KEY")), nil
	})
	// if an error is occurred by parsing the token string then set the err
	if err != nil {
		err = errors.New(errorMessage)
		return
	}
	// last check if the token is valid - it's not return an error
	if !token.Valid {
		err = errors.New(errorMessage)
		return
	}
	return
}