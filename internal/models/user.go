package models

type User struct {
	Id uint64 `gorm:"primary_key"json:"id"binding:"required"`

	// e-mail address from user - will be used as "user-name" // todo: make this uniquely
	EMailAddress string `json:"eMailAddress"binding:"required"`

	// That's the user input by account creation or for verification
	Password string `gorm:"-"json:"password"binding:"required"`

	// That's the persisted password in db
	PasswordHashed string
}
