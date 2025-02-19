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

	config, err := pgx.ParseConfig("")
	if err != nil {
		fmt.Println(err)
	}
	conn, err := pgx.Connect(context.Background(), config.ConnString())
	if err != nil {
		fmt.Println()
	}

	config.Host = os.Getenv("PGHOST")
	config.Port = 5432
	config.User = os.Getenv("PGUSER")
	config.Password = os.Getenv("PGPASSWORD")
	config.Database = os.Getenv("PGDATABASE")
	config.RuntimeParams = map[string]string{
		"sslmode": "disable",
	}
	return conn
}
