package database

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"time"
)

func Config() *pgxpool.Config {
	const defaultMaxConnections = int32(4)
	const defaultMinConncetions = int32(0)
	const defaultMaxConnLifetime = time.Hour
	const defaultMaxConnIdleTime = time.Minute * 30
	const defaultHealthCheckPeriod = time.Minute
	const defaultConnectionTimeout = time.Second * 5

	const DATABASE_URL string = "postgres://postgres:mint@localhost:5432/postgres"

	config, err := pgxpool.ParseConfig(DATABASE_URL)
	if err != nil {
		log.Fatalf("Failed to parse DATABASE_URL: %v", err)
	}

	config.MaxConns = defaultMaxConnections
	config.MinConns = defaultMinConncetions
	config.MaxConnLifetime = defaultMaxConnLifetime
	config.MaxConnIdleTime = defaultMaxConnIdleTime
	config.HealthCheckPeriod = defaultHealthCheckPeriod
	config.ConnConfig.ConnectTimeout = defaultConnectionTimeout

	config.BeforeAcquire = func(ctx context.Context, conn *pgx.Conn) bool {
		log.Println("Before acquiring the connection pool to the database!!")
		return true
	}

	config.AfterRelease = func(conn *pgx.Conn) bool {
		log.Println("After releasing the connection pool to the database!!")
		return true
	}

	config.BeforeClose = func(conn *pgx.Conn) {
		log.Println("Closed the connection pool to the database!!")
	}

	return config
}
