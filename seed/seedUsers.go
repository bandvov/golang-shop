package seeds

import (
	"context"
	"fmt"
	"log"

	"github.com/bandvov/golang-shop/application"
	"github.com/bandvov/golang-shop/domain/users"
	"github.com/bandvov/golang-shop/infrastructure"
	"github.com/jackc/pgx/v4"
)

func Seed(conn *pgx.Conn) {
	userRepo := &infrastructure.PostgresUserRepository{Conn: conn}
	userService := &application.UserService{Repo: userRepo}

	// Seed data
	seedUsers := []*users.User{
		{FirstName: "John", LastName: "Doe", Email: "john.doe@example.com", Password: "password123"},
		{FirstName: "Jane", LastName: "Smith", Email: "jane.smith@example.com", Password: "password123"},
		{FirstName: "Bob", LastName: "Johnson", Email: "bob.johnson@example.com", Password: "password123"},
	}
	for _, user := range seedUsers {
		err := userService.CreateUser(context.Background(), user)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Users seeding completed.")

}
