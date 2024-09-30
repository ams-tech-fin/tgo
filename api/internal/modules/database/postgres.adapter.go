package database

import (
	"fmt"
	"log"
	"tgo/api/internal/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func getDBURL() string {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.GetEnv("DB_HOST", "localhost"),
		config.GetEnv("DB_PORT", "5432"),
		config.GetEnv("DB_USER", "postgres"),
		config.GetEnv("DB_PASSWORD", "secret"),
		config.GetEnv("DB_NAME", "myapp"))

	return connStr
}

func ConnectDB() (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", getDBURL())
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
		return nil, err
	}

	log.Println("Connected to database")
	return db, nil
}
