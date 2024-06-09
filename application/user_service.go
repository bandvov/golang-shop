package application

import "github.com/bandvov/golang-shop/domain/users"

type UserService struct {
	Repo users.UserRepository
}

func NewUserService(repo users.UserRepository) *UserService {
	return &UserService{Repo: repo}
}

func (s *UserService) GetUserS() ([]*users.User, error) {
	return s.Repo.GetUsers()
}
func (s *UserService) GetUserByID(id int) (*users.User, error) {
	return s.Repo.GetByID(id)
}

func (s *UserService) CreateUser(user *users.User) error {
	return s.Repo.Save(user)
}

func (s *UserService) DeleteUser(id int) error {
	return s.Repo.Delete(id)
}

func (s *UserService) UpdateUser(user *users.User) error {
	return s.Repo.Update(user)
}
