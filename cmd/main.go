package main

import (
	"log"
	"net/http"

	"github.com/Domson12/shople_api/configs"
	"github.com/Domson12/shople_api/routes"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	// Load configurations
	configs.LoadConfig()

	// Initialize Chi router
	router := chi.NewRouter()

	// Middleware setup
	router.Use(middleware.Logger)            // Logs requests
	router.Use(middleware.Recoverer)         // Recovers from panics
	router.Use(middleware.Timeout(60 * 1e9)) // Request timeout

	// Register routes
	routes.RegisterRoutes(router)

	// Start the server
	port := configs.GetConfig("APP_PORT")
	log.Printf("Starting server on :%s...", port)
	http.ListenAndServe(":"+port, router)
}
