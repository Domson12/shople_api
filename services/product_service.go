package services

import (
	"github.com/Domson12/shople_api/models"
	"github.com/Domson12/shople_api/repositories"
)

type ProductService struct {
	ProductRepo *repositories.ProductRepo
}

func (service *ProductService) GetProducts() ([]models.Product, error) {
	return service.ProductRepo.GetAllProducts()
}
