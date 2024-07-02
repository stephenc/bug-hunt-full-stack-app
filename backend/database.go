package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

const (
	defaultDbHost     = "localhost"
	defaultDbPort     = 5432
	defaultDbUser     = "postgres"
	defaultDbPassword = "password"
	defaultDbName     = "tododb"
)

func initDB() (*sql.DB, error) {
	// Read the database connection details from environment variables
	dbHost := getEnv("DB_HOST", defaultDbHost)
	dbPort := getEnv("DB_PORT", fmt.Sprintf("%d", defaultDbPort))
	dbUser := getEnv("DB_USER", defaultDbUser)
	dbPassword := getEnv("DB_PASSWORD", defaultDbPassword)
	dbName := getEnv("DB_NAME", defaultDbName)

	// Build the connection string
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	// Open the database connection
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %v", err)
	}

	// Verify the connection
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("error pinging database: %v", err)
	}

	return db, nil
}

// getEnv reads an environment variable or returns a default value if not set
func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}

func createTables() {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		username TEXT UNIQUE NOT NULL,
		password TEXT NOT NULL
	)`)

	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS todos (
		id SERIAL PRIMARY KEY,
		user_id INT NOT NULL,
		text TEXT NOT NULL,
		FOREIGN KEY (user_id) REFERENCES users(id)
	)`)

	if err != nil {
		log.Fatal(err)
	}
}
