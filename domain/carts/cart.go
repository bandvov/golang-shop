package carts

import (
	"errors"
	"time"
)

var (
	ErrCartNotFound = errors.New("cart not found")
)

type CartItem struct {
	CartItemID   int     `json:"cart_item_id"`
	CartID       int     `json:"cart_id"`
	ProductID    int     `json:"product_id"`
	Quantity     int     `json:"quantity"`
	Price        float64 `json:"price"`
	SessionID    string  `json:"session_id"`
	DiscountCode string  `json:"discount_code"`
	Total        float64 `json:"total"`
}

type Cart struct {
	CartID    int        `json:"cart_id"`
	UserID    int        `json:"user_id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	Items     []CartItem `json:"items"`
}

type CartRepository interface {
	GetCarts()([]*Cart,error) 
	Save(cart *Cart) error
	FindByID(id int) (*Cart, error)
	Delete(id int) error
	Update(cart *Cart) error
}
