package users

import (
	"errors"
)

type User struct {
	ID       int
	Name     string `json:"name" validate:"required,min=2,max=100"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

var (
	ErrUserNotFound = errors.New("user not found")
)

type UserRepository interface {
	GetUsers() ([]*User, error)
	GetByID(id int) (*User, error)
	Save(user *User) error
	Delete(id int) error
	Update(user *User) error
}
