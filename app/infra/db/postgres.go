package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

func connectToPostgreSQL() (*sql.DB, error) {
	// Database connection parameters
	username := "your-username"
	password := "your-password"
	hostname := "your-hostname"
	port := "your-port"
	database := "your-database"

	// Create a connection string
	dataSourceName := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable", username, password, hostname, port, database)

	// Open the PostgreSQL database
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}

	// Ping the database to verify the connection
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to PostgreSQL!")
	return db, nil
}

func main() {
	// Connect to PostgreSQL
	db, err := connectToPostgreSQL()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create a table
	createTableSQL := `
        CREATE TABLE IF NOT EXISTS users (
            id SERIAL PRIMARY KEY,
            name VARCHAR(255)
        )
    `
	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}

	// Insert data into the table
	insertSQL := "INSERT INTO users (name) VALUES ($1)"
	_, err = db.Exec(insertSQL, "John Doe")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Data inserted into the table.")
}
