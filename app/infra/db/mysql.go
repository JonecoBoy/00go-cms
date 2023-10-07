package db

import (
	"database/sql"
	"fmt"
	"log"
)

func connectToMySQL() (*sql.DB, error) {
	// Database connection parameters
	username := "your-username"
	password := "your-password"
	hostname := "your-hostname"
	port := "your-port"
	database := "your-database"

	// Create a connection string
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, hostname, port, database)

	// Open the MySQL database
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}

	// Ping the database to verify the connection
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to MySQL!")
	return db, nil
}

func main() {
	// Connect to MySQL
	db, err := connectToMySQL()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create a table
	createTableSQL := `
        CREATE TABLE IF NOT EXISTS users (
            id INT AUTO_INCREMENT PRIMARY KEY,
            name VARCHAR(255)
        )
    `
	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}

	// Insert data into the table
	insertSQL := "INSERT INTO users (name) VALUES (?)"
	_, err = db.Exec(insertSQL, "John Doe")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Data inserted into the table.")
}
