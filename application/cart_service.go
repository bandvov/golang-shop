package application

import "github.com/bandvov/golang-shop/domain/carts"

type CartService struct {
	Repo carts.CartRepository
}

func NewCartService(repo carts.CartRepository) *CartService {
	return &CartService{Repo: repo}
}

func (s *CartService) GetCarts() ([]*carts.Cart, error) {
	return s.Repo.GetCarts()
}

func (s *CartService) AddToCart(cart *carts.Cart) error {
	return s.Repo.Save(cart)
}

func (s *CartService) GetCartByID(id int) (*carts.Cart, error) {
	return s.Repo.FindByID(id)
}

func (s *CartService) RemoveFromCart(id int) error {
	return s.Repo.Delete(id)
}

func (s *CartService) UpdateCart(cart *carts.Cart) error {
	return s.Repo.Update(cart)
}
