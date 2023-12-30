// router/router.go
package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pnlyzer/handlers"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()

	// Define routes
	router.HandleFunc("/", handlers.HomeHandler).Methods("GET")

	return router
}
