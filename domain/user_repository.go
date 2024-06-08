
package domain

type UserRepository interface {
    FindAll() ([]*User, error)
    FindByID(id int) (*User, error)
    Save(user *User) error
    Update(user *User) error
    Delete(id int) error
}
