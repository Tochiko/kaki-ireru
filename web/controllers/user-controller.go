package controllers

import (
	"encoding/base64"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/kataras/iris/core/errors"
	"kaki-ireru/internal/models"
	"kaki-ireru/internal/provider"
	"net/http"
	"os"
	"strings"
	"time"
)

// create a new user by sending eMail and password via json
func CreateUser (c *gin.Context) {
	// bind the json input in an user object
	var user models.User
	c.BindJSON(&user)
	// create the bounded user
	err := provider.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	// if everything worked well respond with status 204 - No Content
	c.Status(http.StatusNoContent)
}

// basic http authorization needs username and password from users
// a bearer token is returned by success
func Login (c *gin.Context) {
	c.Header("WWW-Authenticate", "Basic http://localhost:3000/login")
	// extract the credentials for login from authorization header
	credentials, err := extractCredentials(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}
	// verify the credentials from authorization header with persisted user
	id, verification := provider.VerifyUser(credentials[0], credentials[1])
	if !verification {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "username or password are wrong"})
		return
	}
	// create a bearer token, set the claims and receive it as binary
	binaryToken, err := createToken(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "something's gone wrong please try it again"})
		return
	}
	// if everything worked well remove the www-authenticate header and write the token as binary
	c.Header("WWW-Authenticate", "")
	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Write(binaryToken)
}

/**
extract the credentials from a basic http authorization
the result contains user eMails on index zero and the corresponding password on index one
 */
func extractCredentials (c *gin.Context) (credentials []string, err error) {
	errorMessage := "authorization header must be set correctly"
	// 2 because the value is "basic [base64 encoded credentials]"
	authValue := strings.SplitN(c.GetHeader("Authorization"), " ", 2)
	if len(authValue) != 2 {
		err = errors.New(errorMessage)
		return
	}
	// 1 because that's the index of the credentials [basic, credentials]
	decAuthValue, err := base64.StdEncoding.DecodeString(authValue[1])
	if err != nil {
		return
	}
	// 2 because the credentials looks like username:password
	authPair := strings.SplitN(string(decAuthValue), ":", 2)
	if len(authPair) != 2 {
		err = errors.New(errorMessage)
		return
	}
	credentials = authPair
	return
}

/**
Create a HS256 signed jwt token with the following claims
id - the id form an user; expiration - from the token
 */
func createToken (id int) (binaryToken []byte, err error) {
	// create a new jwt.Token with claims - user id and expiration is set
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": id,
		"expiration": time.Now().Add(time.Hour * 24).Unix(),
	})
	// sign the created token with private key
	tokenString, err := token.SignedString([]byte(os.Getenv("HS_PRIVATE_KEY")))
	if err != nil {
		return
	}
	// if everything worked well return the token as binary
	binaryToken = []byte(tokenString)
	return
}
