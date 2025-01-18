package repositories

import (
	"database/sql"

	"github.com/Domson12/shople_api/models"
)

type ProductRepo struct {
	DB *sql.DB
}

func (repo *ProductRepo) GetAllProducts() ([]models.Product, error) {
	products := []models.Product{}
	rows, err := repo.DB.Query("SELECT id, name, description, price, stock, image_url FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var product models.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Stock, &product.ImageURL); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}
