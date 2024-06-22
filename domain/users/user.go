package users

import (
	"errors"
	"time"
)

type User struct {
	ID        int
	FirstName string    `json:"first_name,omitempty" validate:"required,min=2,max=100"`
	LastName  string    `json:"last_name,omitempty" validate:"required,min=2,max=100"`
	Email     string    `json:"email,omitempty" validate:"required,email"`
	Password  string    `json:"password,omitempty" validate:"required,min=8"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	Role      string    `json:"role,omitempty" validate:"required,oneof=admin user"`
	Status    string    `json:"status,omitempty" validate:"required,oneof=active inactive"`
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
