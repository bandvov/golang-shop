package seeds

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/bandvov/golang-shop/application"
	"github.com/bandvov/golang-shop/domain/users"
	"github.com/bandvov/golang-shop/infrastructure"
)

func Seed(db *sql.DB) {
	userRepo := &infrastructure.PostgresUserRepository{DB: db}
	userService := &application.UserService{Repo: userRepo}

	// Seed data
	seedUsers := []*users.User{
		{FirstName: "John", LastName: "Doe", Email: "john.doe@example.com", Password: "password123"},
		{FirstName: "Jane", LastName: "Smith", Email: "jane.smith@example.com", Password: "password123"},
		{FirstName: "Bob", LastName: "Johnson", Email: "bob.johnson@example.com", Password: "password123"},
	}

	for _, user := range seedUsers {
		err := userService.CreateUser(user)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Database setup and seeding completed.")

}
