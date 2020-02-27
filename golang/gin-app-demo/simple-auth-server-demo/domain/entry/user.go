package entry

import (
	"time"
)

type User struct {
	ID                uint `gorm:"primary_key"`
	Email             string
	EncryptedPassword string
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         *time.Time
}
