package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/fashion-reels/database"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {

	// Create a new connection pool
	pool, err := pgxpool.NewWithConfig(context.Background(), database.Config())
	if err != nil {
		log.Fatalf("Failed to create the connection pool: %v", err)
	}

	connection, err := pool.Acquire(context.Background())
	if err != nil {
		log.Fatalf("Failed to acquire the connection pool: %v", err)
	}
	defer connection.Release()

	err = connection.Ping(context.Background())
	if err != nil {
		log.Fatalf("Failed to ping the connection pool: %v", err)
	}

	fmt.Println("Successfully connected to the database!!")

	// Server Code
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server listening on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
