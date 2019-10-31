package model

// 地域
type Area struct {
	ID                int    // ID
	Name              string // 地域名
	CustomerCompanies []CustomerCompany
}
