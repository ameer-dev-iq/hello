package models

import (
	"fmt"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uint   `json:"id" gorm:"primary_key" `
	Username string `json:"username" gorm:"not null" validate:"required"`
	Password string `json:"password" gorm:"not null" validate:"required"`
}

func Validate(user *User) error {
	if user.Username == "" && len(user.Username) < 3 {
		return fmt.Errorf("username is required")
	}
	if user.Password == "" && len(user.Password) < 8 {
		return fmt.Errorf("password is required")
	}
	return nil
}
