// application/user_service.go
package application

import "github.com/bandvov/golang-shop/domain"


type UserService struct {
    UserRepo domain.UserRepository
}

func (s *UserService) GetUsers() ([]*domain.User, error) {
    return s.UserRepo.FindAll()
}

func (s *UserService) GetUserByID(id int) (*domain.User, error) {
    return s.UserRepo.FindByID(id)
}

func (s *UserService) CreateUser(user *domain.User) error {
    return s.UserRepo.Save(user)
}

func (s *UserService) UpdateUser(user *domain.User) error {
    return s.UserRepo.Update(user)
}

func (s *UserService) DeleteUser(id int) error {
    return s.UserRepo.Delete(id)
}
