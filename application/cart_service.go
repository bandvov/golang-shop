package application

import (
	"context"

	"github.com/bandvov/golang-shop/domain/carts"
)

type CartService struct {
	Repo carts.CartRepository
}

func NewCartService(repo carts.CartRepository) *CartService {
	return &CartService{Repo: repo}
}

func (s *CartService) GetCarts(ctx context.Context) ([]*carts.Cart, error) {
	return s.Repo.GetCarts(ctx)
}

func (s *CartService) AddToCart(ctx context.Context, cart *carts.Cart) error {
	return s.Repo.Save(ctx, cart)
}

func (s *CartService) GetCartByID(ctx context.Context, id int) (*carts.Cart, error) {
	return s.Repo.FindByID(ctx, id)
}

func (s *CartService) RemoveFromCart(ctx context.Context, id int) error {
	return s.Repo.Delete(ctx, id)
}

func (s *CartService) UpdateCart(ctx context.Context, cart *carts.Cart) error {
	return s.Repo.Update(ctx, cart)
}
