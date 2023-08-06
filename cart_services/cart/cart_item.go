package cart

type CartItem struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func NewCartItem(name string, price float64) *CartItem {
	return &CartItem{
		Name:  name,
		Price: price,
	}
}
