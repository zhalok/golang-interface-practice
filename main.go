package main

import (
	"fmt"
	"time"
)

type Product struct {
	Name     string
	Price    float64
	Category string
	Quantity int
}

type ProductRepository interface {
	getProductsByCategory(category string) ([]Product, error)
}

type ProductService struct {
	productRepository ProductRepository
}

func (ps *ProductService) getAvaialbleProducts(category string) ([]Product, error) {
	allProducts, error := ps.productRepository.getProductsByCategory(category)
	if error != nil {
		return []Product{}, error
	}
	availableProducts := []Product{}
	for _, product := range allProducts {
		if product.Quantity > 0 {
			availableProducts = append(availableProducts, product)
		}

	}
	return availableProducts, nil
}

type productRepositoryImpl struct {
}

func (pr *productRepositoryImpl) getProductsByCategory(category string) ([]Product, error) {
	fmt.Println("Connecting to database...")
	time.Sleep(time.Second * 2)

	return []Product{
		{
			Name:     "Laptop",
			Price:    1000.00,
			Category: "Electronics",
			Quantity: 10,
		},
		{
			Name:     "TV",
			Price:    400.00,
			Category: "Electronics",
			Quantity: 5,
		},
	}, nil
}

func main() {
	productService := ProductService{
		productRepository: &productRepositoryImpl{}}

	availableProducts, error := productService.getAvaialbleProducts("Electronics")
	if error != nil {
		panic(error)
	}
	for _, product := range availableProducts {
		println(product.Name)
	}
}
