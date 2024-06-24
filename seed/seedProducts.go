package seeds

import (
	"context"
	"fmt"

	"github.com/bandvov/golang-shop/application"
	"github.com/bandvov/golang-shop/domain/products"
	"github.com/bandvov/golang-shop/infrastructure"
	"github.com/jackc/pgx/v4"
)

func SeedProducts(conn *pgx.Conn) {
	productRepo := infrastructure.PostgresProductRepository{Conn: conn}
	productService := application.ProductService{Repo: &productRepo}

	// Sample products to seed
	products := []*products.Product{
		{Name: "Product 1", Description: "Description 1", Price: 10.99, SKU: "SKU001", StockQuantity: 100, Category: "Category A", Brand: "Brand A", ImageURL: "https://example.com/image1.jpg"},
		{Name: "Product 2", Description: "Description 2", Price: 20.99, SKU: "SKU002", StockQuantity: 150, Category: "Category B", Brand: "Brand B", ImageURL: "https://example.com/image2.jpg"},
		{Name: "Product 3", Description: "Description 3", Price: 30.99, SKU: "SKU003", StockQuantity: 200, Category: "Category C", Brand: "Brand C", ImageURL: "https://example.com/image3.jpg"},
	}

	// Insert each product into the database
	for _, product := range products {
		err := productService.CreateProduct(context.Background(), product)
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("Product seeding completed.")
}
