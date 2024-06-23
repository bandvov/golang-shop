package products

import (
	"errors"
	"time"
)

type Product struct {
	ID            int       `json:"id,omitempty" db:"id"`
	Name          string    `json:"name,omitempty" db:"name"`
	Description   string    `json:"description,omitempty" db:"description"`
	Price         float64   `json:"price,omitempty" db:"price"`
	SKU           string    `json:"sku,omitempty" db:"sku"`
	StockQuantity int       `json:"stock_quantity,omitempty" db:"stock_quantity"`
	Category      string    `json:"category,omitempty" db:"category"`
	Brand         string    `json:"brand,omitempty" db:"brand"`
	ImageURL      string    `json:"image_url,omitempty" db:"image_url"`
	CreatedAt     time.Time `json:"created_at,omitempty" db:"created_at,omitempty"`
	UpdatedAt     time.Time `json:"updated_at,omitempty" db:"updated_at"`
	IsActive      bool      `json:"is_active,omitempty" db:"is_active"`
}

var (
	ErrProductNotFound = errors.New("product not found")
)

type ProductRepository interface {
	GetProducts() ([]*Product, error)
	GetByID(id int) (*Product, error)
	Save(product *Product) error
	Delete(id int) error
	Update(product *Product) error
}
