package model

// 個人顧客
type BusinessCustomer struct {
	*Customer     // 会員顧客を継承
	Capital   int // 資本金
}

func NewBusinessCustomer(id int, name, email string, capital int) *BusinessCustomer {
	return &BusinessCustomer{
		Customer: &Customer{
			ID:    id,
			Name:  name,
			Email: email,
		},
		Capital: capital,
	}
}
