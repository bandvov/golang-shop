package users

import (
	"context"
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
	Role      string    `json:"role,omitempty" validate:"oneof=admin user"`
	Status    string    `json:"status,omitempty" validate:"oneof=active inactive"`
}

var (
	ErrUserNotFound = errors.New("user not found")
)

type UserRepository interface {
	GetUsers(ctx context.Context) ([]*User, error)
	GetByID(ctx context.Context, id int) (*User, error)
	GetByEmail(ctx context.Context, email string) (*User, error)
	Save(ctx context.Context, user *User) error
	Delete(ctx context.Context, id int) error
	Update(ctx context.Context, user *User) error
}
