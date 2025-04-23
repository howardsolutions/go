package store

import (
	"database/sql"
	"embed"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
	"github.com/pressly/goose/v3"
)

func Open() (*sql.DB, error) {
	// Try to load environment variables, but don't error if .env file doesn't exist
	_ = godotenv.Load()

	// Get database URL from environment variable
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		// Use default connection string if DATABASE_URL is not set
		dbURL = "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
		fmt.Println("No DATABASE_URL found, using default:", dbURL)
	}

	// Parse the config
	config, err := pgx.ParseConfig(dbURL)
	if err != nil {
		return nil, fmt.Errorf("unable to parse database URL: %w", err)
	}

	// Create a *sql.DB using the pgx driver
	db := stdlib.OpenDB(*config)

	// Test the connection
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %w", err)
	}

	fmt.Println("Connected to Database....")
	return db, nil
}

func MigrateFS(db *sql.DB, migrationsFS embed.FS, dir string) error {
	err := goose.SetDialect("postgres")
	if err != nil {
		return fmt.Errorf("set dialect: %w", err)
	}

	goose.SetBaseFS(migrationsFS)
	defer func() {
		goose.SetBaseFS(nil)
	}()

	err = goose.Up(db, dir)
	if err != nil {
		return fmt.Errorf("goose up: %w", err)
	}

	return nil
}
