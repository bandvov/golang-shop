package infrastructure

import (
	"context"
	"database/sql"
	"time"

	"github.com/bandvov/golang-shop/domain/users"
	"github.com/jackc/pgx/v4"
	"golang.org/x/crypto/bcrypt"
)

type PostgresUserRepository struct {
	Conn *pgx.Conn
}

func NewPostgresUserRepository(conn *pgx.Conn) *PostgresUserRepository {
	return &PostgresUserRepository{Conn: conn}
}

func (r *PostgresUserRepository) GetUsers(ctx context.Context) ([]*users.User, error) {

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var u []*users.User
	query := "SELECT id, first_name, last_name, email, role, status FROM users"

	rows, err := r.Conn.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var user users.User
		if err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Role, &user.Status); err != nil {
			return nil, err
		}
		u = append(u, &user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return u, nil

}

func (r *PostgresUserRepository) GetByID(ctx context.Context, id int) (*users.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var user users.User
	query := "SELECT id, first_name,last_name, email, role, status, created_at FROM users WHERE id=$1"
	err := r.Conn.QueryRow(ctx, query, id).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Role, &user.Status, &user.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, users.ErrUserNotFound
		}
		return nil, err
	}
	return &user, nil
}

func (r *PostgresUserRepository) GetByEmail(ctx context.Context, email string) (*users.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var user users.User

	query := `SELECT id, email, password, created_at, updated_at FROM users WHERE email=$1`

	err := r.Conn.QueryRow(ctx, query, email).Scan(&user.ID, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err == pgx.ErrNoRows {
		return nil, users.ErrUserNotFound
	}
	return &user, err
}

func (r *PostgresUserRepository) Save(ctx context.Context, user *users.User) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	query := "INSERT INTO users (first_name,last_name, email, password) VALUES ($1, $2, $3, $4)"
	_, err = r.Conn.Exec(ctx, query, user.FirstName, user.LastName, user.Email, string(hashedPassword))
	return err
}

func (r *PostgresUserRepository) Delete(ctx context.Context, id int) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	query := "DELETE FROM users WHERE id=$1"
	_, err := r.Conn.Exec(ctx, query, id)
	return err
}

func (r *PostgresUserRepository) Update(ctx context.Context, user *users.User) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	query := "UPDATE users SET first_name=$1,first_name=$2, email=$3, WHERE id=$5"
	_, err := r.Conn.Exec(ctx, query, user.FirstName, user.LastName, user.Email, user.ID)
	return err
}
