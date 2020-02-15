package model

// 業種
type Industry struct {
	ID                int    // ID
	Name              string // 業種名
	CustomerCompanies []CustomerCompany
}
