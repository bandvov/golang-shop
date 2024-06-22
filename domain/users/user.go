package users

import (
	"errors"
	"time"
)

type User struct {
	ID        int
	FirstName string    `json:"first_name" validate:"required,min=2,max=100"`
	LastName  string    `json:"last_name" validate:"required,min=2,max=100"`
	Email     string    `json:"email" validate:"required,email"`
	Password  string    `json:"password" validate:"required,min=8"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Role      string    `json:"role" validate:"required,oneof=admin user"`
	Status    string    `json:"status" validate:"required,oneof=active inactive"`
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
