package db

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var Pool *pgxpool.Pool

func ConnectDB() {
	var err error

	// Get DB URL env
	databaseUrl := os.Getenv("DATABASE_URL")

	// Create connection
	Pool, err = pgxpool.New(context.Background(), databaseUrl)
	if err != nil {
		log.Fatalln("Error connecting to DB", err)
	}

	// Ping database
	err = Pool.Ping(context.Background())
	if err != nil {
		log.Fatalln("Cannot ping DB", err)
	}

	log.Println("Connected to DB")
}
