package controllers

import (
	"encoding/base64"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"kaki-ireru/internal/models"
	"kaki-ireru/internal/provider"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func CreateUser (c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	err := provider.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	} else {
		c.Status(http.StatusNoContent)
		return
	}
}

func Login (c *gin.Context) {
	c.Header("WWW-Authenticate", "Basic http://localhost:3000/login")
	authValue := strings.SplitN(c.GetHeader("Authorization"), " ", 2)
	if len(authValue) != 2 { // 2 because the value is "basic [base64 encoded credentials]"
		c.Status(http.StatusUnauthorized)
		return
	}

	decAuthValue, err := base64.StdEncoding.DecodeString(authValue[1]) // 1 because that's the index of the credentials
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	authPair := strings.SplitN(string(decAuthValue), ":", 2) // 2 because the credentials looks like username:password
	if len(authPair) != 2 {
		c.Status(http.StatusUnauthorized)
		return
	}
	id, verification := provider.VerifyUser(authPair[0], authPair[1])
	if verification {
		c.Header("WWW-Authenticate", "")
		token := jwt.New(jwt.SigningMethodHS256)
		claims := token.Claims.(jwt.MapClaims)
		claims["id"] = strconv.Itoa(id)
		claims["eMail"] = authPair[0]
		claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
		tokenString, err := token.SignedString([]byte(os.Getenv("HS_PRIVATE_KEY")))
		if err != nil {
			fmt.Println(err.Error())
		}
		c.Writer.WriteHeader(http.StatusOK)
		c.Writer.Write([]byte(tokenString))
		return
	} else {
		c.Status(http.StatusUnauthorized)
		return
	}
}