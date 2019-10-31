package model

// 顧客企業
type CustomerCompany struct {
	ID          int        // ID
	Name        string     // 社名
	Location    string     // 所在地
	Capital     int        // 資本金
	CreditLimit int        // 与信限度額
	AreaID      int        // 地域 ID
	Industries  []Industry // 業種
}

var i int = 1

func NewCustomerCompany(name, location string, capital, creditLimit int) *CustomerCompany {
	return &CustomerCompany{
		ID:          i,
		Name:        name,
		Location:    location,
		Capital:     capital,
		CreditLimit: creditLimit,
	}
}
