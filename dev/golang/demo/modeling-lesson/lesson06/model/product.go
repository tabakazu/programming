package model

// 商品
type Product struct {
	ID    int    // ID
	Name  string // 名称
	Price int    // 値段
}

func NewProduct(id int, name string, price int) *Product {
	return &Product{
		ID:    id,
		Name:  name,
		Price: price,
	}
}
