// infrastructure/postgres_user_repository.go
package infrastructure

import (
	"database/sql"
	"log"

	"github.com/bandvov/golang-shop/domain"
)

type PostgresUserRepository struct {
	DB *sql.DB
}

func (r *PostgresUserRepository) FindAll() ([]*domain.User, error) {
	rows, err := r.DB.Query("SELECT id, username, email FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var users []*domain.User
	for rows.Next() {
		var user domain.User
		err := rows.Scan(&user.ID, &user.Username, &user.Email)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, &user)
	}
	return users, nil
}

func (r *PostgresUserRepository) FindByID(id int) (*domain.User, error) {
	var user domain.User
	err := r.DB.QueryRow("SELECT id, username, email FROM users WHERE id = $1", id).Scan(&user.ID, &user.Username, &user.Email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *PostgresUserRepository) Save(user *domain.User) error {
	_, err := r.DB.Exec("INSERT INTO users (username, email) VALUES ($1, $2)", user.Username, user.Email)
	if err != nil {
		return err
	}
	return nil
}

func (r *PostgresUserRepository) Update(user *domain.User) error {
	_, err := r.DB.Exec("UPDATE users SET username=$1, email=$2 WHERE id=$3", user.Username, user.Email, user.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *PostgresUserRepository) Delete(id int) error {
	_, err := r.DB.Exec("DELETE FROM users WHERE id=$1", id)
	if err != nil {
		return err
	}
	return nil
}
