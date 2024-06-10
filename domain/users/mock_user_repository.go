package users


type MockUserRepository struct {
    GetByIDFunc func(id string) (*User, error)
    SaveFunc    func(user *User) error
    DeleteFunc  func(id string) error
    UpdateFunc  func(user *User) error
}

func (m *MockUserRepository) GetByID(id string) (*User, error) {
    if m.GetByIDFunc != nil {
        return m.GetByIDFunc(id)
    }
    return nil, ErrUserNotFound
}

func (m *MockUserRepository) Save(user *User) error {
    if m.SaveFunc != nil {
        return m.SaveFunc(user)
    }
    return nil
}

func (m *MockUserRepository) Delete(id string) error {
    if m.DeleteFunc != nil {
        return m.DeleteFunc(id)
    }
    return nil
}

func (m *MockUserRepository) Update(user *User) error {
    if m.UpdateFunc != nil {
        return m.UpdateFunc(user)
    }
    return nil
}
