package service

import (
	"go-practice/model"
	"go-practice/repository"
)

type ProductService struct {
	ProductRepository repository.ProductRepository
}

func (ps *ProductService) GetAvaialbleProducts(category string) ([]model.Product, error) {
	allProducts, error := ps.ProductRepository.GetProductsByCategory(category)
	if error != nil {
		return []model.Product{}, error
	}
	availableProducts := []model.Product{}
	for _, product := range allProducts {
		if product.Quantity > 0 {
			availableProducts = append(availableProducts, product)
		}

	}
	return availableProducts, nil
}
