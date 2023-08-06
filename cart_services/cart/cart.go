package cart

import (
	"encoding/json"
	"fmt"
)

type Cart struct {
	items []*CartItem
}

func NewCart() *Cart {
	return &Cart{
		items: []*CartItem{},
	}
}

func (c *Cart) AddItem(item *CartItem) {
	c.items = append(c.items, item)
}

func (c *Cart) PrintCartItems() {
	fmt.Println("Productos en el carrito:")
	for _, item := range c.items {
		fmt.Printf("- %s: %.2f\n", item.Name, item.Price)
	}
}

func (c *Cart) GetCartItems() ([]*CartItem, error) {
	return c.items, nil
}

func (c *Cart) GetCartItemsJSON() ([]byte, error) {
	data := make([]map[string]interface{}, len(c.items))
	for i, item := range c.items {
		data[i] = map[string]interface{}{
			"name":  item.Name,
			"price": item.Price,
		}
	}

	return json.Marshal(data)
}
