package users

import "context"

type MockUserRepository struct {
	GetUsersFunc   func(ctx context.Context) ([]*User, error)
	GetByIDFunc    func(ctx context.Context, id int) (*User, error)
	GetByEmailFunc func(ctx context.Context, email string) (*User, error)
	SaveFunc       func(ctx context.Context, user *User) error
	DeleteFunc     func(ctx context.Context, id int) error
	UpdateFunc     func(ctx context.Context, user *User) error
}

func (m *MockUserRepository) GetUsers(ctx context.Context) ([]*User, error) {
	if m.GetUsersFunc != nil {
		return m.GetUsersFunc(ctx)
	}
	return nil, ErrUserNotFound
}
func (m *MockUserRepository) GetByID(ctx context.Context, id int) (*User, error) {
	if m.GetByIDFunc != nil {
		return m.GetByIDFunc(ctx, id)
	}
	return nil, ErrUserNotFound
}

func (m *MockUserRepository) Save(ctx context.Context, user *User) error {
	if m.SaveFunc != nil {
		return m.SaveFunc(ctx, user)
	}
	return nil
}

func (m *MockUserRepository) Delete(ctx context.Context, id int) error {
	if m.DeleteFunc != nil {
		return m.DeleteFunc(ctx, id)
	}
	return nil
}

func (m *MockUserRepository) Update(ctx context.Context, user *User) error {
	if m.UpdateFunc != nil {
		return m.UpdateFunc(ctx, user)
	}
	return nil
}
func (m *MockUserRepository) GetByEmail(ctx context.Context, email string) (*User, error) {
	if m.GetByEmailFunc != nil {
		return m.GetByEmailFunc(ctx, email)
	}
	return nil, nil
}
