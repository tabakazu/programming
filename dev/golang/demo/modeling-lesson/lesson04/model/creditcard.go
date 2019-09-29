package model

import "time"

// クレジットカード
type Creditcard struct {
	ID         int       // ID
	CustomerID int       // 会員顧客 ID
	CardNumber string    // カード番号
	Expiration time.Time // 有効期限
	Nominee    string    // 名義
}

func NewCreditcard(id, customerID int, cardNumber string, expiration time.Time, nominee string) *Creditcard {
	return &Creditcard{
		ID:         id,
		CustomerID: customerID,
		CardNumber: cardNumber,
		Expiration: expiration,
		Nominee:    nominee,
	}
}
