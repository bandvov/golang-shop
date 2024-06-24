package application

import (
	"context"

	"github.com/bandvov/golang-shop/domain/products"
)

type ProductService struct {
	Repo products.ProductRepository
}

func NewProductService(repo products.ProductRepository) *ProductService {
	return &ProductService{Repo: repo}
}

func (s *ProductService) GetProducts(ctx context.Context) ([]*products.Product, error) {
	return s.Repo.GetProducts(ctx)
}
func (s *ProductService) GetProductByID(ctx context.Context, id int) (*products.Product, error) {
	return s.Repo.GetByID(ctx, id)
}

func (s *ProductService) CreateProduct(ctx context.Context, product *products.Product) error {
	return s.Repo.Save(ctx, product)
}

func (s *ProductService) DeleteProduct(ctx context.Context, id int) error {
	return s.Repo.Delete(ctx, id)
}

func (s *ProductService) UpdateProduct(ctx context.Context, product *products.Product) error {
	return s.Repo.Update(ctx, product)
}
