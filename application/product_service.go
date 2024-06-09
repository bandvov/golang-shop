package application

import "github.com/bandvov/golang-shop/domain/products"

type ProductService struct {
	Repo products.ProductRepository
}

func NewProductService(repo products.ProductRepository) *ProductService {
	return &ProductService{Repo: repo}
}

func (s *ProductService) GetProducts() ([]*products.Product, error) {
	return s.Repo.GetProducts()
}
func (s *ProductService) GetProductByID(id int) (*products.Product, error) {
	return s.Repo.GetByID(id)
}

func (s *ProductService) CreateProduct(product *products.Product) error {
	return s.Repo.Save(product)
}

func (s *ProductService) DeleteProduct(id int) error {
	return s.Repo.Delete(id)
}

func (s *ProductService) UpdateProduct(product *products.Product) error {
	return s.Repo.Update(product)
}
