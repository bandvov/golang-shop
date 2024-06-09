package products

import "errors"

type Product struct {
    ID          int
    Name        string
    Description string
    Price       float64
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
