package entity

type User struct {
	Model
	Email          string `gorm:"unique;not null"`
	Password       string `gorm:"-" json:"-"`
	PasswordDigest string `gorm:"not null" json:"-"`
}
