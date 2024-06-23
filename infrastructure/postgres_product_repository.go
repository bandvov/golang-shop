package infrastructure

import (
	"database/sql"

	"github.com/bandvov/golang-shop/domain/products"
	_ "github.com/lib/pq"
)

type PostgresProductRepository struct {
	DB *sql.DB
}

func NewPostgresProductRepository(db *sql.DB) *PostgresProductRepository {
	return &PostgresProductRepository{DB: db}
}
func (r *PostgresProductRepository) GetProducts() ([]*products.Product, error) {
	var p []*products.Product
	query := "SELECT id, name, description FROM products"

	rows, err := r.DB.Query(query)
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

func (r *PostgresProductRepository) GetByID(id int) (*products.Product, error) {
	var product products.Product
	query := "SELECT id, name, description, price FROM products WHERE id=$1"
	err := r.DB.QueryRow(query, id).Scan(&product.ID, &product.Name, &product.Description, &product.Price)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, products.ErrProductNotFound
		}
		return nil, err
	}
	return &product, nil
}

func (r *PostgresProductRepository) Save(product *products.Product) error {
	query := "INSERT INTO products (name, description, price, sku, stock_quantity,category,brand ) VALUES ($1, $2, $3, $4, $5, $6, $7)"
	_, err := r.DB.Exec(query, product.Name, product.Description, product.Price, product.SKU, product.StockQuantity, product.Category, product.Brand)
	return err
}

func (r *PostgresProductRepository) Delete(id int) error {
	query := "DELETE FROM products WHERE id=$1"
	_, err := r.DB.Exec(query, id)
	return err
}

func (r *PostgresProductRepository) Update(product *products.Product) error {
	query := "UPDATE products SET name=$1, description=$2, price=$3 WHERE id=$4"
	_, err := r.DB.Exec(query, product.Name, product.Description, product.Price, product.ID)
	return err
}
