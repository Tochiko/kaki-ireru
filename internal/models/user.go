package models

type User struct {
	Id int `gorm:"primary_key"json:"id"binding:"required"`

	// e-mail address from user - will be used as "user-name"
	EMailAddress string `gorm:"unique"json:"eMailAddress"binding:"required"`

	// That's the user input by account creation or for verification
	Password string `gorm:"-"json:"password"binding:"required"`

	// That's the persisted password in db
	PasswordHashed string


}