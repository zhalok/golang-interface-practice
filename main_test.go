package main

import (
	"fmt"
	"go-practice/model"
	"go-practice/repository"
	"go-practice/service"
	"testing"
)

type ProductRepositoryMockImplSuccess struct {
}

func (pr *ProductRepositoryMockImplSuccess) GetProductsByCategory(category string) ([]model.Product, error) {
	fmt.Println("Mocking product repository...")
	return []model.Product{
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

type ProductRepositoryMockImplFailed struct {
}

func (pr *ProductRepositoryMockImplFailed) GetProductsByCategory(category string) ([]model.Product, error) {
	fmt.Println("Mocking product repository...")
	return []model.Product{
		{
			Name:     "Laptop",
			Price:    1000.00,
			Category: "Electronics",
			Quantity: 0,
		},
		{
			Name:     "TV",
			Price:    400.00,
			Category: "Electronics",
			Quantity: 5,
		},
	}, nil
}

func TestMain(t *testing.T) {

	cases := []struct {
		productRepository repository.ProductRepository
		expected          int
	}{
		{
			productRepository: &ProductRepositoryMockImplSuccess{},

			expected: 2,
		},
		{
			productRepository: &ProductRepositoryMockImplFailed{},

			expected: 2,
		},
	}
	for _, c := range cases {
		productService := service.ProductService{
			ProductRepository: c.productRepository}
		availableProducts, error := productService.GetAvaialbleProducts("Electronics")
		if error != nil {
			panic(error)
		}
		if len(availableProducts) != c.expected {
			t.Errorf("Expected %v but got %v", c.expected, len(availableProducts))
		}
	}
}
