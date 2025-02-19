package lib

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func DB() *pgx.Conn {
	godotenv.Load()

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("PGHOST"),
		os.Getenv("PGPORT"),
		os.Getenv("PGUSER"),
		os.Getenv("PGPASSWORD"),
		os.Getenv("PGDATABASE"),
	)

	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		fmt.Println("Error connecting to database:", err)
	}

	return conn
}
