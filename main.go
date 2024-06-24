package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/bandvov/golang-shop/application"
	"github.com/bandvov/golang-shop/infrastructure"
	"github.com/bandvov/golang-shop/interfaces"
	"github.com/jackc/pgx/v4"
)

func main() {
	if err := loadEnv(".env"); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	dbName := os.Getenv("POSTGRES_DATABASE_NAME")
	dbHost := os.Getenv("POSTGRES_DATABASE_HOST")
	dbUser := os.Getenv("POSTGRES_DATABASE_USER")
	dbPort := os.Getenv("POSTGRES_DATABASE_PORT")
	dbUserPassword := os.Getenv("POSTGRES_DATABASE_USER_PASSWORD")

	connStr := fmt.Sprintf("postgres://%v:%v@%v:%v/%v", dbUser, dbUserPassword,dbHost, dbPort, dbName)

	// Connect to PostgreSQL
	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer conn.Close(context.Background())

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}

	// seeds.Seed(db)
	// Initialize repositories
	userRepo := &infrastructure.PostgresUserRepository{Conn: conn}
	productRepo := &infrastructure.PostgresProductRepository{Conn: conn}
	cartRepo := &infrastructure.PostgresCartRepository{Conn: conn}

	// Initialize services
	userService := &application.UserService{Repo: userRepo}
	productService := &application.ProductService{Repo: productRepo}
	cartService := &application.CartService{Repo: cartRepo}
	// seeds.SeedProducts(db)
	// Initialize handlers
	userHandler := &interfaces.UserHandler{UserService: userService}
	productHandler := &interfaces.ProductHandler{ProductService: productService}
	cartHandler := &interfaces.CartHandler{CartService: cartService}

	// Create a new ServeMux for the entire application
	mux := http.NewServeMux()
	mux.Handle("GET /users", interfaces.LoggerMiddleware(userHandler.GetUsers))
	mux.Handle("POST /users", interfaces.LoggerMiddleware(userHandler.CreateUser))
	mux.Handle("PUT /users", interfaces.LoggerMiddleware(userHandler.UpdateUser))
	mux.Handle("GET /users/{id}", interfaces.LoggerMiddleware(userHandler.GetUserByID))
	mux.Handle("DELETE /users/{id}", interfaces.LoggerMiddleware(userHandler.DeleteUser))

	mux.Handle("GET /products", interfaces.LoggerMiddleware(productHandler.GetProducts))
	mux.Handle("POST /products", interfaces.LoggerMiddleware(productHandler.CreateProduct))
	mux.Handle("PUT /products", interfaces.LoggerMiddleware(productHandler.UpdateProduct))
	mux.Handle("GET /products/{id}", interfaces.LoggerMiddleware(productHandler.GetProduct))
	mux.Handle("DELETE /products/{id}", interfaces.LoggerMiddleware(productHandler.DeleteProduct))

	mux.Handle("GET /carts", interfaces.LoggerMiddleware(cartHandler.GetCarts))
	mux.Handle("POST /carts", interfaces.LoggerMiddleware(cartHandler.AddToCart))
	mux.Handle("PUT /carts", interfaces.LoggerMiddleware(cartHandler.UpdateCart))
	mux.Handle("GET /carts/{id}", interfaces.LoggerMiddleware(cartHandler.GetCart))
	mux.Handle("DELETE /carts/{id}", interfaces.LoggerMiddleware(cartHandler.RemoveFromCart))

	log.Printf("Starting server on %v\n", PORT)
	http.ListenAndServe(fmt.Sprintf(":%v", PORT), mux)

}

func loadEnv(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			log.Printf("Invalid line in %s: %s", filename, line)
			continue
		}
		key, value := parts[0], parts[1]
		os.Setenv(key, value)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
