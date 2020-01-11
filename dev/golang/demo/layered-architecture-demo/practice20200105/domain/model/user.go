package model

type User struct {
	Model
	Email string `gorm:"column:email" json:"email"`
}
