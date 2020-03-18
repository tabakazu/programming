package valueobject

type Address struct {
	Street  string `gorm:"column:street"`
	City    string `gorm:"column:city"`
	Country string `gorm:"column:country"`
}
