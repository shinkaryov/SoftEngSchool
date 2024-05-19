package connect_to_db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq" // PostgreSQL driver
)

// ConnectDB establishes a connection to the PostgreSQL database and returns the connection
func ConnectDB() (*sql.DB, error) {

	// Access environment variables
	postgresUser := os.Getenv("POSTGRES_USER")
	postgresDB := os.Getenv("POSTGRES_DB")
	postgresPassword := os.Getenv("POSTGRES_PASSWORD")
	// postgresHost should be the name of the PostgreSQL service in the docker-compose file
	postgresHost := "db"
	postgresPort := "5432"

	// Create the connection string
	connStr := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%s sslmode=disable",
		postgresUser, postgresDB, postgresPassword, postgresHost, postgresPort)

	// Open the database connection
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}

	// Check the connection
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("error connecting to the database: %w", err)
	}

	fmt.Println("Successfully connected to the database")
	return db, nil
}
