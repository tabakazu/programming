package entity

type User struct {
	Model
	Email          string `gorm:"unique;not null"`
	Password       string `gorm:"-"`
	PasswordDigest string `gorm:"not null"`
}
