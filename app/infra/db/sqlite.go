package db

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func connectToSQLite() (*sql.DB, error) {
	// Open the SQLite database
	db, err := sql.Open("sqlite3", "your-database-path")
	if err != nil {
		return nil, err
	}

	// Ping the database to verify the connection
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to SQLite!")
	return db, nil
}

func main() {
	// Connect to SQLite
	db, err := connectToSQLite()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create a table
	createTableSQL := `
        CREATE TABLE IF NOT EXISTS users (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            name TEXT
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
