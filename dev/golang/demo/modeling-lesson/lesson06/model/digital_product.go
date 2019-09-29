package model

// 商品
type DigitalProduct struct {
	*Product
	Url string // URL
}

func NewDigitalProduct(id int, name string, price int, url string) *DigitalProduct {
	return &DigitalProduct{
		Product: &Product{
			ID:    id,
			Name:  name,
			Price: price,
		},
		Url: url,
	}
}
