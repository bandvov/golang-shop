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

	connStr := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v  sslmode=verify-full sslrootcert=%v", dbHost, dbPort, dbUser, dbUserPassword, dbName,"")

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

	// Initialize services
	userService := &application.UserService{UserRepo: userRepo}

	// Initialize handlers
	userHandler := &interfaces.UserHandler{UserService: userService}

	// Create a new ServeMux for the entire application
	mux := http.NewServeMux()
	mux.HandleFunc("/users", interfaces.LoggerMiddleware(userHandler).ServeHTTP)

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
