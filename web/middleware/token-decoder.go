package middleware

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/kataras/iris/core/errors"
	"net/http"
	"os"
	"strings"
)

type Jwks struct {
	Keys []JSONWebKeys `json:"keys"`
}

type JSONWebKeys struct {
	Kty string `json:"kty"`
	Kid string `json:"kid"`
	Use string `json:"use"`
	N string `json:"n"`
	E string `json:"e"`
	X5c []string `json:"x5c"`
}

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
	errorMessage := "the received token is not allowed"
	// parse the token string into a jwt token
	// the inner function is initially checking if tokens signing method can be cast with SigningMethodHMAC
	// if the cast has success so tokens method is ok and the private key is returned from inner function
	/*token, err = jwt.Parse(tokenString, func (t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf(errorMessage)
		}
		return []byte(os.Getenv("HS_PRIVATE_KEY")), nil
	})*/

	token, err = jwt.Parse(tokenString, keyFunc)

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

func keyFunc(token *jwt.Token) (interface{}, error) {
	aud := os.Getenv("AUTH0_AUDIENCE")
	checkAud := token.Claims.(jwt.MapClaims).VerifyAudience(aud, false)
	if !checkAud {
		return token, errors.New("Invalid audience.")
	}
	// Verify 'iss' claim
	iss := "https://" + os.Getenv("AUTH0_DOMAIN") + "/"
	checkIss := token.Claims.(jwt.MapClaims).VerifyIssuer(iss, false)
	if !checkIss {
		return token, errors.New("Invalid issuer.")
	}

	cert, err := getPemCert(token)
	if err != nil {
		panic(err.Error())
	}

	result, _ := jwt.ParseRSAPublicKeyFromPEM([]byte(cert))
	return result, nil
}

func getPemCert(token *jwt.Token) (string, error) {
	cert := ""
	resp, err := http.Get("https://" + os.Getenv("AUTH0_DOMAIN") + "/.well-known/jwks.json")

	if err != nil {
		return cert, err
	}
	defer resp.Body.Close()

	var jwks = Jwks{}
	err = json.NewDecoder(resp.Body).Decode(&jwks)

	if err != nil {
		return cert, err
	}

	for k := range jwks.Keys {
		if token.Header["kid"] == jwks.Keys[k].Kid {
			cert = "-----BEGIN CERTIFICATE-----\n" + jwks.Keys[k].X5c[0] + "\n-----END CERTIFICATE-----"
		}
	}

	if cert == "" {
		err := errors.New("Unable to find appropriate key.")
		return cert, err
	}

	return cert, nil
}