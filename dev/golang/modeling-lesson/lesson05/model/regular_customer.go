package model

import "time"

// 個人顧客
type RegularCustomer struct {
	*Customer     // 会員顧客を継承
	Sex       int // 性別
	Birthday  time.Time
}

func NewRegularCustomer(id int, name, email string, sex int, birthday time.Time) *RegularCustomer {
	return &RegularCustomer{
		Customer: &Customer{
			ID:    id,
			Name:  name,
			Email: email,
		},
		Sex:      sex,
		Birthday: birthday,
	}
}
