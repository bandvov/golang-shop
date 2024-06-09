package users

import (
	"errors"
)

type User struct {
	ID       int
	Name     string
	Email    string
	Password string
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
