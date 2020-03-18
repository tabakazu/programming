package valueobject

type CustomerName struct {
	Value string `gorm:"column:name"`
}
