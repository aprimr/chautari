package db

import (
	"context"
	"os"

	"github.com/aprimr/chautari/utils"
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
		utils.LogFatal("Error connecting to DB", err)
	}

	// Ping database
	err = Pool.Ping(context.Background())
	if err != nil {
		utils.LogFatal("Cannot ping DB", err)
	}

	utils.LogInfo("Connected to DB")
}
