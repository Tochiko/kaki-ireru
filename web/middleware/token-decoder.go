package middleware

import (
	"encoding/json"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strings"
)

type JwkSet struct {
	Keys []JsonWebKey `json:"keys"`
}

type JsonWebKey struct {
	Kty string `json:"kty"`
	Kid string `json:"kid"`
	Use string `json:"use"`
	N string `json:"n"`
	E string `json:"e"`
	X5c []string `json:"x5c"`
}

// Middleware function for gin gonic
// Bearer tokens are decoded and the sub claim is set as userId to context
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

	// extract the user id from claims it's not more needed at this time
	claims := token.Claims.(jwt.MapClaims)
	userId := claims["sub"]
	if userId == "" {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "Unknown identity"})
		return
	}

	// if everything worked well remove the www-authenticate header and set token claims to context
	c.Header("WWW-Authenticate", "")
	c.Set("userId", userId)
}

/**
extract token strings from http authorization header
in case of errors a descriptive message is returned
 */
func extractTokenString (c *gin.Context) (tokenString string, err error) {
	errorMessage := "you have to set the authorization header with bearer token correctly to use this service"
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
	// parse the token string into a jwt token by keyFunc
	token, err = jwt.Parse(tokenString, keyFunc)

	// if an error is occurred by parsing the token string then set the err
	if err != nil {
		return
	}
	// last check if the token is valid - it's not return an error
	if !token.Valid {
		err = errors.New("the received token is not valid")
		return
	}
	return
}

// function to parse a token string into jwt.Token
// first the audience and the issuer are checked
// in the following the pem certificate is get to parse the token with RSA public key
// if everything worked well the parsed token is returned as interface{}
func keyFunc(token *jwt.Token) (interface{}, error) {
	// verify the audience with the given one in the process variables
	aud := os.Getenv("AUTH0_AUDIENCE")
	audOk := token.Claims.(jwt.MapClaims).VerifyAudience(aud, false)
	if !audOk {
		return token, errors.New("audience is not valid")
	}

	// verify the issuer by building the URI with given auth0 domain from process variables
	issuer := "https://" + os.Getenv("AUTH0_DOMAIN") + "/"
	issuerOk := token.Claims.(jwt.MapClaims).VerifyIssuer(issuer, false)
	if !issuerOk {
		return token, errors.New("issuer is not valid")
	}

	// get the pem certificate to parse the token afterwards
	cert, err := getPemCert(token)
	if err != nil {
		return token, err
	}

	// parse the token with pem certificate
	result, err := jwt.ParseRSAPublicKeyFromPEM([]byte(cert))
	if err != nil {
		return token, nil
	}

	// if everything worked well return the parsed token as interface{}
	return result, nil
}

// get the pem certificate from auth0 identity provider by request the jwks
func getPemCert(token *jwt.Token) (string, error) {
	cert := ""

	// get request on jwks from auth0 by given auth0 domain in process variables
	resp, err := http.Get("https://" + os.Getenv("AUTH0_DOMAIN") + "/.well-known/jwks.json")
	if err != nil {
		return cert, err
	}
	defer resp.Body.Close()

	// create an instance of JwkSet and decode the resp body on it
	var jwks = JwkSet{}
	err = json.NewDecoder(resp.Body).Decode(&jwks)
	if err != nil {
		return cert, err
	}

	// iterate the jwkSet and find the matching jwk by comparing the kid property with the value from the token
	// if a jwk is found then fill the cert with X5c property and return
	for k := range jwks.Keys {
		if token.Header["kid"] == jwks.Keys[k].Kid {
			cert = "-----BEGIN CERTIFICATE-----\n" + jwks.Keys[k].X5c[0] + "\n-----END CERTIFICATE-----"
			return cert, nil
		}
	}

	// no matching jwk was in the jwkSet? then create an error and return
	err = errors.New("no matching jwk found")
	return cert, err
}