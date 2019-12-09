package model

import (
	"time"
)

type User struct {
	ID             uint   `gorm:"primary_key"`
	Email          string `gorm:"unique_index"`
	Password       string `gorm:"-"`
	PasswordDigest string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      *time.Time
}

func NewUser(email, password string) *User {
	return &User{
		Email:    email,
		Password: password,
	}
}

// func (u *User) encryptPassword() error {
// 	if u.Password != "" {
// 		hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
// 		if err != nil {
// 			return err
// 		}
// 		u.PasswordDigest = string(hash)
// 	}
// 	return nil
// }
