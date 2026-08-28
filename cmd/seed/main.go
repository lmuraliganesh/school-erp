package main

import (
	"context"
	"fmt"
	"log"

	"school-erp/database"
	"school-erp/utils"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, relying on system environment variables")
	}

	// Connect to the database
	database.ConnectDB()

	// Admin details
	name := "School Admin"
	email := "admin@mangoschool.in"
	plainPassword := "ChangeMe123!"
	role := "admin"

	// Hash the password
	hashedPassword, err := utils.HashPassword(plainPassword)
	if err != nil {
		log.Fatalf("Failed to hash password: %v", err)
	}

	// Insert the admin into the database
	query := `INSERT INTO users (name, email, password_hash, role) VALUES ($1, $2, $3, $4)`
	_, err = database.DB.Exec(context.Background(), query, name, email, hashedPassword, role)
	if err != nil {
		log.Fatalf("Failed to insert admin: %v", err)
	}

	fmt.Println("Admin created:", email)
}