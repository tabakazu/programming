package valueobject

type CustomerID struct {
	Value uint `gorm:"column:customer_id"`
}
