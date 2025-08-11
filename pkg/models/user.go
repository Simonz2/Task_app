package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// for next version implemt an active field
// to check if the user has been activated by the
// organization
type User struct {
	gorm.Model
	Username  string `json:"username" gorm:"uniqueIndex;not null"`
	Password  string `json:"password" gorm:"not null"`
	Activated bool   `json:"activated" gorm:"default:false"`
}

// hashes and sets the user's password
func (u *User) SetPassword(password string) error {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedBytes)
	return nil
}

// CheckPassword compares given password with stored hash
func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
