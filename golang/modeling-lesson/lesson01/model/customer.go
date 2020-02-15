package model

// 会員顧客
type Customer struct {
	ID int // ID
}

func NewCustomer(id int) *Customer {
	return &Customer{id}
}
