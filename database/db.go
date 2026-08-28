package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

// ConnectDB initializes the PostgreSQL connection pool
func ConnectDB() {
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		log.Fatal("DATABASE_URL is not set in environment variables")
	}

	var err error
	DB, err = pgxpool.New(context.Background(), databaseURL)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	// Test the connection
	err = DB.Ping(context.Background())
	if err != nil {
		log.Fatalf("Database ping failed: %v\n", err)
	}

	fmt.Println("Successfully connected to PostgreSQL database!")
}

// InitUserTable creates the users table if it doesn't already exist
func InitUserTable() {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		email VARCHAR(100) UNIQUE NOT NULL,
		password_hash VARCHAR(255) NOT NULL,
		role VARCHAR(20) NOT NULL CHECK (role IN ('admin', 'teacher', 'student')),
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	`

	_, err := DB.Exec(context.Background(), query)
	if err != nil {
		log.Fatalf("Failed to create users table: %v\n", err)
	}

	fmt.Println("Users table checked/created successfully!")
}
