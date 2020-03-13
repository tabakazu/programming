package domain

import (
	"time"
)

type User struct {
	ID             uint `gorm:"primary_key"`
	Email          string
	PasswordDigest string
	APIToken       string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      *time.Time
}
