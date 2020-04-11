package entity

import "github.com/tabakazu/gortfolio/domain/value"

type User struct {
	Model
	Email          value.Email `gorm:"embedded"`
	Password       string      `gorm:"-" json:"-"`
	PasswordDigest string      `gorm:"not null" json:"-"`
}
