package application

import (
	"context"

	"github.com/bandvov/golang-shop/domain/users"
)

type UserService struct {
	Repo users.UserRepository
}

func NewUserService(repo users.UserRepository) *UserService {
	return &UserService{Repo: repo}
}

func (s *UserService) GetUsers(ctx context.Context) ([]*users.User, error) {
	return s.Repo.GetUsers(ctx)
}

func (s *UserService) GetUserByID(ctx context.Context, id int) (*users.User, error) {
	return s.Repo.GetByID(ctx, id)
}

func (s *UserService) GetUserByEmail(ctx context.Context, email string) (*users.User, error) {
	return s.Repo.GetByEmail(ctx, email)
}

func (s *UserService) CreateUser(ctx context.Context, user *users.User) error {
	return s.Repo.Save(ctx, user)
}

func (s *UserService) DeleteUser(ctx context.Context, id int) error {
	return s.Repo.Delete(ctx, id)
}

func (s *UserService) UpdateUser(ctx context.Context, user *users.User) error {
	return s.Repo.Update(ctx, user)
}
