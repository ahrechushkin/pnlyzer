package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"pnlyzer/models"
	"pnlyzer/router"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	port := getPort()
	db, err := initializeDB()

	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}

	// Create the server with routes
	srv := createServer(port, router.NewRouter(db))

	// Run the server in a goroutine
	go func() {
		log.Printf("Server is running on port %s...\n", port)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Error starting server: %v\n", err)
		}
	}()

	// Set up signal handling for graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	// Block until a signal is received
	<-stop
	log.Println("Shutting down server...")

	// Create a context with a timeout to allow outstanding requests to finish
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Shutdown the server gracefully
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Error shutting down server: %v\n", err)
	}

	log.Println("Server gracefully stopped")
}

func getPort() string {
	port := os.Getenv("PNLYZER_PORT")

	if port == "" {
		port = "8080"
	}

	return port
}

func createServer(port string, router *mux.Router) *http.Server {
	return &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}
}

func initializeDB() (*gorm.DB, error) {
	dcs := getDBConnectionString()
	db, err := gorm.Open(postgres.Open(dcs), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	log.Println("Connected to the database")

	// Run auto-migration to create tables
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		return nil, err
	}

	return db, nil

}

func getDBConnectionString() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_PORT"),
		os.Getenv("DATABASE_USER"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_NAME"),
	)
}
