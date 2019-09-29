package model

// 商品
type RealProduct struct {
	*Product
	StockCount int // 在庫数
}

func NewRealProduct(id int, name string, price, stockCount int) *RealProduct {
	return &RealProduct{
		Product: &Product{
			ID:    id,
			Name:  name,
			Price: price,
		},
		StockCount: stockCount,
	}
}
