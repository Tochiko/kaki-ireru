package provider

import (
	"github.com/kataras/iris/core/errors"
	"golang.org/x/crypto/bcrypt"
	"kaki-ireru/internal/models"
)

// create a new user
func CreateUser (user *models.User) (err error) {
	// generate a hash from password
	hashedPassword, e := generatePasswordHash(user.Password)
	if e != nil {
		err = errors.New("the password is not allowed")
		return
	}
	// set the generated password hash to corresponding user as string and create it
	user.PasswordHashed = string(hashedPassword)
	if e := Db.Create(&user).Error; e != nil {
		err = errors.New("the e-mail address is already in use")
	}
	// if everything worked well the returned error is nil
	return
}

// verify a user by credentials
// returns the id from user and true if the verification was successful
func VerifyUser (eMail string, password string) (id int, verification bool) {
	// create a user object with the credentials
	var user models.User
	// initially check if a user with given e-mail address is existing
	if err := Db.Where("e_mail_address = ?", eMail).First(&user).Error; err != nil {
		verification = false
		return
	}
	// verify the user password from input with the persisted password hash from db
	verification = compareHashedPassword(user.PasswordHashed, password)
	// if everything worked well then return the user id and the result from verification
	id = user.Id
	return
}

/**
generate a hashedPassword from password string and return it as binary
 */
func generatePasswordHash(password string) (hashedPassword []byte, err error) {
	hashedPassword, err = bcrypt.GenerateFromPassword([]byte(password), 8)
	return
}

/**
compare a hashedPassword with a human readable password
returns true when the comparison is ok
 */
func compareHashedPassword (hashedPassword string, password string) bool {
	// compare the hashedPassword from db with password from input
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return false
	}
	// if there is no error then the verification is successfully
	return true
}