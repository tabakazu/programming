package value

type Email struct {
	Value string `gorm:"column:email;unique;not null"`
}
