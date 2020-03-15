package valueobject

type PostName struct {
	Value string `gorm:"column:name"`
}
