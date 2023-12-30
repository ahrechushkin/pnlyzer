package main

import (
	"context"
	"errors"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"pnlyzer/router"
	"time"
)

func main() {
	port := getPort()

	// Create the server with routes
	srv := createServer(port, router.NewRouter())

	// Run the server in a goroutine
	go func() {
		log.Printf("Server is running on port %s...\n", port)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Error starting server: %v", err)
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
		log.Fatalf("Error shutting down server: %v", err)
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
