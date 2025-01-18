package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Domson12/shople_api/services"
)

type ProductController struct {
	ProductService *services.ProductService
}

func (ctrl *ProductController) GetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := ctrl.ProductService.GetProducts()
	if err != nil {
		http.Error(w, "Failed to fetch products", http.StatusInternalServerError)
		return
	}

	// Write response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}
