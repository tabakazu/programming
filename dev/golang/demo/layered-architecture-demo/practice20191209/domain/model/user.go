package model

import (
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
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

func (u *User) encryptPassword() error {
	if u.Password == "" {
		return fmt.Errorf("u.Password is blank")
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.PasswordDigest = string(hash)
	return nil
}

func UserAuthenticate(email, password string) (*User, error) {
	// ユーザの取得を模倣
	u := &User{
		Email:    email,
		Password: password,
	}
	if err := u.encryptPassword(); err != nil {
		return nil, err
	}

	hash := u.PasswordDigest
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return nil, err
	}
	return u, nil
}
