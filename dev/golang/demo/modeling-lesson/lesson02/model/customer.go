package model

// 会員顧客
type Customer struct {
	Name  string
	Email string
}

func NewCustomer(name, email string) *Customer {
	return &Customer{name, email}
}
