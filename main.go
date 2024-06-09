package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/bandvov/golang-shop/application"
	"github.com/bandvov/golang-shop/infrastructure"
	"github.com/bandvov/golang-shop/interfaces"
	_ "github.com/lib/pq"
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

	connStr := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v  sslmode=verify-full sslrootcert=%v", dbHost, dbPort, dbUser, dbUserPassword, dbName, "")

	// Connect to PostgreSQL
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}

	// Initialize repositories
	userRepo := &infrastructure.PostgresUserRepository{DB: db}
	productRepo := &infrastructure.PostgresProductRepository{DB: db}

	// Initialize services
	userService := &application.UserService{Repo: userRepo}
	productService := &application.ProductService{Repo: productRepo}

	// Initialize handlers
	userHandler := &interfaces.UserHandler{UserService: userService}
	productHandler := &interfaces.ProductHandler{ProductService: productService}

	// Create a new ServeMux for the entire application
	mux := http.NewServeMux()
	mux.Handle("GET /users", interfaces.LoggerMiddleware(userHandler.GetUsers))
	mux.Handle("PUT /users", interfaces.LoggerMiddleware(userHandler.UpdateUser))
	mux.Handle("GET /users/{id}", interfaces.LoggerMiddleware(userHandler.GetUserByID))
	mux.Handle("DELETE /users/{id}", interfaces.LoggerMiddleware(userHandler.DeleteUser))

	mux.Handle("GET /products", interfaces.LoggerMiddleware(productHandler.GetProducts))
	mux.Handle("PUT /products", interfaces.LoggerMiddleware(productHandler.UpdateProduct))
	mux.Handle("GET /products/{id}", interfaces.LoggerMiddleware(productHandler.GetProduct))
	mux.Handle("DELETE /products/{id}", interfaces.LoggerMiddleware(productHandler.DeleteProduct))

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
