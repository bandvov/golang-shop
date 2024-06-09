package infrastructure

import (
	"database/sql"

	"github.com/bandvov/golang-shop/domain/users"
	_ "github.com/lib/pq"
)

type PostgresUserRepository struct {
	DB *sql.DB
}

func NewPostgresUserRepository(db *sql.DB) *PostgresUserRepository {
	return &PostgresUserRepository{DB: db}
}

func (r *PostgresUserRepository) GetUsers() ([]*users.User, error) {
	var u []*users.User
	query := "SELECT id, name, email, password FROM users"

	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var user *users.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			return nil, err
		}
		u = append(u, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return u, nil

}

func (r *PostgresUserRepository) GetByID(id int) (*users.User, error) {
	var user users.User
	query := "SELECT id, name, email, password FROM users WHERE id=$1"
	err := r.DB.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, users.ErrUserNotFound
		}
		return nil, err
	}
	return &user, nil
}

func (r *PostgresUserRepository) Save(user *users.User) error {
	query := "INSERT INTO users (id, name, email, password) VALUES ($1, $2, $3, $4)"
	_, err := r.DB.Exec(query, user.ID, user.Name, user.Email, user.Password)
	return err
}

func (r *PostgresUserRepository) Delete(id int) error {
	query := "DELETE FROM users WHERE id=$1"
	_, err := r.DB.Exec(query, id)
	return err
}

func (r *PostgresUserRepository) Update(user *users.User) error {
	query := "UPDATE users SET name=$1, email=$2, password=$3 WHERE id=$4"
	_, err := r.DB.Exec(query, user.Name, user.Email, user.Password, user.ID)
	return err
}
