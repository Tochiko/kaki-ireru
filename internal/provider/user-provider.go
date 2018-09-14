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

func VerifyUser (eMail string, password string) (id int, verification bool) {
	var user models.User
	if err := Db.Where("e_mail_address = ?", eMail).First(&user).Error; err != nil {
		verification = false
		return
	}
	verification = compareHashedPassword(&user, password)
	id = user.Id
	return
}

func hashPassword (user *models.User) (*models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
	user.PasswordHashed = string(hashedPassword)
	return user, err
}

func compareHashedPassword (user * models.User, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHashed), []byte(password))
	if err != nil {
		return false
	} else {
		return true
	}
}