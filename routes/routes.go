package routes

import (
	"database/sql"

	"github.com/Domson12/shople_api/controllers"
	"github.com/Domson12/shople_api/repositories"
	"github.com/Domson12/shople_api/services"

	"github.com/go-chi/chi/v5"
	_ "github.com/lib/pq" // PostgreSQL driver
)

func RegisterRoutes(router *chi.Mux) {
	// Connect to database
	db, err := sql.Open("postgres", "postgresql://user:password@localhost:5432/shopdb?sslmode=disable")
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}

	// Setup dependencies
	productRepo := &repositories.ProductRepo{DB: db}
	productService := &services.ProductService{ProductRepo: productRepo}
	productController := &controllers.ProductController{ProductService: productService}

	// Product routes
	router.Route("/products", func(r chi.Router) {
		r.Get("/", productController.GetProducts) // GET /products
	})
}
