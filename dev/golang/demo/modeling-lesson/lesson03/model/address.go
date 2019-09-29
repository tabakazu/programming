package model

// 配送先
type Address struct {
	ID            int    // ID
	CustomerID    int    // 会員顧客 ID
	Name          string // 宛名
	StreetAddress string // 住所
}

func NewAddress(id, customerID int, name, streetAddress string) *Address {
	return &Address{
		ID:            id,
		CustomerID:    customerID,
		Name:          name,
		StreetAddress: streetAddress,
	}
}
