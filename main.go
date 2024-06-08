package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	dbName := os.Getenv("POSTGRES_DATABASE_NAME")
	dbHost := os.Getenv("POSTGRES_DATABASE_HOST")
	dbUser := os.Getenv("POSTGRES_DATABASE_USER")
	dbPort := os.Getenv("POSTGRES_DATABASE_PORT")
	dbUserPassword := os.Getenv("POSTGRES_DATABASE_USER_PASSWORD")

	connStr := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=disable", dbHost, dbPort, dbUser, dbUserPassword, dbName)

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
	log.Printf("Starting server on %v\n", PORT)
	http.ListenAndServe(fmt.Sprintf(":%v", PORT), nil)

}
