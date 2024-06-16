package infrastructure

import (
	"database/sql"

	"github.com/bandvov/golang-shop/domain/carts"
	_ "github.com/lib/pq"
)

type PostgresCartRepository struct {
	DB *sql.DB
}

func NewPostgresCartRepository(db *sql.DB) carts.CartRepository {
	return &PostgresCartRepository{DB: db}
}

func (r *PostgresCartRepository) Save(cart *carts.Cart) error {
	tx, err := r.DB.Begin()
	if err != nil {
		return err
	}

	query := `INSERT INTO carts (user_id, created_at, updated_at) VALUES ($1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP) RETURNING cart_id`
	err = tx.QueryRow(query, cart.UserID).Scan(&cart.CartID)
	if err != nil {
		tx.Rollback()
		return err
	}

	for _, item := range cart.Items {
		query := `INSERT INTO cart_items (cart_id, product_id, quantity, price, session_id, discount_code, total) VALUES ($1, $2, $3, $4, $5, $6, $3 * $4)`
		_, err := tx.Exec(query, cart.CartID, item.ProductID, item.Quantity, item.Price, item.SessionID, item.DiscountCode)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}

func (r *PostgresCartRepository) GetCarts() ([]*carts.Cart, error) {
	var p []*carts.Cart
	query := "SELECT * FROM carts"

	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var product *carts.Cart
		if err := rows.Scan(&product); err != nil {
			return nil, err
		}
		p = append(p, product)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return p, nil
}

func (r *PostgresCartRepository) FindByID(id int) (*carts.Cart, error) {
	cart := &carts.Cart{}
	query := `SELECT cart_id, user_id, created_at, updated_at FROM carts WHERE cart_id = $1`
	err := r.DB.QueryRow(query, id).Scan(&cart.CartID, &cart.UserID, &cart.CreatedAt, &cart.UpdatedAt)
	if err != nil {
		return nil, err
	}

	query = `SELECT cart_item_id, cart_id, product_id, quantity, price, session_id, discount_code, total FROM cart_items WHERE cart_id = $1`
	rows, err := r.DB.Query(query, cart.CartID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		item := carts.CartItem{}
		err := rows.Scan(&item.CartItemID, &item.CartID, &item.ProductID, &item.Quantity, &item.Price, &item.SessionID, &item.DiscountCode, &item.Total)
		if err != nil {
			return nil, err
		}
		cart.Items = append(cart.Items, item)
	}

	return cart, nil
}

func (r *PostgresCartRepository) Delete(id int) error {
	tx, err := r.DB.Begin()
	if err != nil {
		return err
	}

	query := `DELETE FROM cart_items WHERE cart_id = $1`
	_, err = tx.Exec(query, id)
	if err != nil {
		tx.Rollback()
		return err
	}

	query = `DELETE FROM carts WHERE cart_id = $1`
	_, err = tx.Exec(query, id)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (r *PostgresCartRepository) Update(cart *carts.Cart) error {
	tx, err := r.DB.Begin()
	if err != nil {
		return err
	}

	query := `UPDATE carts SET user_id = $1, updated_at = CURRENT_TIMESTAMP WHERE cart_id = $2`
	_, err = tx.Exec(query, cart.UserID, cart.CartID)
	if err != nil {
		tx.Rollback()
		return err
	}

	for _, item := range cart.Items {
		query := `UPDATE cart_items SET quantity = $1, price = $2, session_id = $3, discount_code = $4, total = $1 * $2 WHERE cart_item_id = $5`
		_, err := tx.Exec(query, item.Quantity, item.Price, item.SessionID, item.DiscountCode, item.CartItemID)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}
