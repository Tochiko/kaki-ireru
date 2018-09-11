package provider

import (
	"golang.org/x/crypto/bcrypt"
	"kaki-ireru/internal/models"
)

func CreateUser (user *models.User) (error) {
	user, err := hashPassword(user)
	if err != nil {
		return err
	} else {
		err := Db.Create(user).Error
		return err
	}
}

func VerifyUser (user *models.User) (verification bool) {
	Db.First(&user, user.Id)
	verification = compareHashedPassword(user)
	return
}

func hashPassword (user *models.User) (*models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
	user.PasswordHashed = string(hashedPassword)
	return user, err
}

func compareHashedPassword (user * models.User) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHashed), []byte(user.Password))
	if err != nil {
		return false
	} else {
		return true
	}
}