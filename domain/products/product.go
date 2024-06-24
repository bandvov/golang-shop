package products

import (
	"context"
	"errors"
	"time"
)

type Product struct {
	ID            int       `json:"id,omitempty" db:"id,omitempty"`
	Name          string    `json:"name,omitempty" db:"name"`
	Description   string    `json:"description,omitempty" db:"description"`
	Price         float64   `json:"price,omitempty" db:"price"`
	SKU           string    `json:"sku,omitempty" db:"sku"`
	StockQuantity int       `json:"stock_quantity,omitempty" db:"stock_quantity"`
	Category      string    `json:"category,omitempty" db:"category"`
	Brand         string    `json:"brand,omitempty" db:"brand"`
	ImageURL      string    `json:"image_url,omitempty" db:"image_url,omitempty"`
	CreatedAt     time.Time `json:"created_at,omitempty" db:"created_at,omitempty"`
	UpdatedAt     time.Time `json:"updated_at,omitempty" db:"updated_at,omitempty"`
	IsActive      bool      `json:"is_active,omitempty" db:"is_active,omitempty"`
}

var (
	ErrProductNotFound = errors.New("product not found")
)

type ProductRepository interface {
	GetProducts(ctx context.Context) ([]*Product, error)
	GetByID(ctx context.Context, id int) (*Product, error)
	Save(ctx context.Context, product *Product) error
	Delete(ctx context.Context, id int) error
	Update(ctx context.Context, product *Product) error
}
