// router/router.go
package router

import (
	"github.com/gorilla/mux"
	"pnlyzer/handlers"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()

	// Define routes
	router.HandleFunc("/", handlers.HomeHandler).Methods("GET")

	return router
}
