package infrastructure

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/bandvov/golang-shop/domain/products"
	"github.com/jackc/pgx/v4"
)

type PostgresProductRepository struct {
	Conn *pgx.Conn
}

func NewPostgresProductRepository(conn *pgx.Conn) *PostgresProductRepository {
	return &PostgresProductRepository{Conn: conn}
}
func (r *PostgresProductRepository) GetProducts(ctx context.Context) ([]*products.Product, error) {
	var p []*products.Product
	query := "SELECT id, name, description FROM products"

	rows, err := r.Conn.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var product products.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Description); err != nil {
			return nil, err
		}
		p = append(p, &product)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return p, nil
}

func (r *PostgresProductRepository) GetByID(ctx context.Context, id int) (*products.Product, error) {
	var product products.Product
	query := "SELECT id, name, description, price FROM products WHERE id=$1"
	err := r.Conn.QueryRow(ctx, query, id).Scan(&product.ID, &product.Name, &product.Description, &product.Price)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, products.ErrProductNotFound
		}
		return nil, err
	}
	return &product, nil
}

func (r *PostgresProductRepository) Save(ctx context.Context, product *products.Product) error {
	fmt.Printf("product: %+v\n", product)
	query := "INSERT INTO products (name, description, price, sku, stock_quantity,category,brand ) VALUES ($1, $2, $3, $4, $5, $6, $7)"
	_, err := r.Conn.Exec(ctx, query, product.Name, product.Description, product.Price, product.SKU, product.StockQuantity, product.Category, product.Brand)
	fmt.Errorf("err:%v\n", err)
	return err
}

func (r *PostgresProductRepository) Delete(ctx context.Context, id int) error {
	query := "DELETE FROM products WHERE id=$1"
	_, err := r.Conn.Exec(ctx, query, id)
	return err
}

func (r *PostgresProductRepository) Update(ctx context.Context, product *products.Product) error {
	query := "UPDATE products SET name=$1, description=$2, price=$3 WHERE id=$4"
	_, err := r.Conn.Exec(ctx, query, product.Name, product.Description, product.Price, product.ID)
	return err
}
