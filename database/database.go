package database

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func ConnectDB() {

	databaseUrl := os.Getenv("DATABASE_URL")

	dbpool, err := pgxpool.New(context.Background(), databaseUrl)

	if err != nil {
		log.Fatal("Unable to connect to database:", err)
	}

	err = dbpool.Ping(context.Background())

	if err != nil {
		log.Fatal("Database ping failed:", err)
	}

	log.Println("Connected to PostgreSQL!")

	DB = dbpool
}
