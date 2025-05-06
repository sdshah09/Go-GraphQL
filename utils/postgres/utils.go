package postgres

import (
	"database/sql"
	"fmt"
	"log"

	"my-graphql-server/config"

	_ "github.com/lib/pq"
)

func ConnectDB(cfg *config.Config) (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHOST,
		cfg.DBPort,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %v", err)
	}

	// Test the connection
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("error connecting to the database: %v", err)
	}

	log.Println("Successfully connected to database!")
	return db, nil
}

func CreatePatientsTable(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS patients (
		id SERIAL PRIMARY KEY,
		first_name VARCHAR(255) NOT NULL,
		last_name VARCHAR(255)
	);`

	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("failed to create patients table: %w", err)
	}
	return nil
}
