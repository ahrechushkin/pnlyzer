package router

import (
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"log"
	"pnlyzer/handlers"
	"pnlyzer/models"
)

func NewRouter(db *gorm.DB) *mux.Router {
	router := mux.NewRouter()

	if db == nil {
		log.Fatalln("Error using database...")
	}

	// users routes
	userRepo := models.NewUserRepositoryGorm(db)
	userHandler := handlers.NewUserHandler(*userRepo)
	router.HandleFunc("/signup", userHandler.Signup).Methods("POST")
	router.HandleFunc("/signin", userHandler.SignIn).Methods("POST")

	// system routes
	router.HandleFunc("/", handlers.HomeHandler).Methods("GET")

	return router
}
