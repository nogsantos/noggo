package domain

import (
	"log"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Base
	Name string `json:"name" gorm:"type:varchar(255)"`
	Email string `json:"email" gorm:"type:varchar(255);unique_index"`
	Password string `json:"-" gorm:"type:varchar(255)"`
	Token string `json:"token" gorm:"type:varchar(255);unique_index"`
}


func NewUser() *User {
	return &User{}
}

// Prepare is a function used to preparate the user data before persist it
func (user *User) Prepare() error {

	password, encript_error := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if encript_error != nil {
		log.Fatalf("Error on password encription %v", encript_error.Error())
	}

	user.Password = string(password)
	user.Token = uuid.NewV4().String()

	validate_error := user.validate()

	return validate_error
}

func (user *User) validate() error {
	return nil
}