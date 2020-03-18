package entity

import "github.com/tabakazu/domain-modeling-demo/domain/model/valueobject"

type Customer struct {
	Name    valueobject.CustomerName `gorm:"embedded"`
	Address valueobject.Address      `gorm:"embedded"`
}

func NewCustomer(name valueobject.CustomerName, address valueobject.Address) *Customer {
	return &Customer{
		Name:    name,
		Address: address,
	}
}
