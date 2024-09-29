package main

import (
	"fmt"
	"testing"
)

type ProductRepositoryMockImplSuccess struct {
}

func (pr *ProductRepositoryMockImplSuccess) getProductsByCategory(category string) ([]Product, error) {
	fmt.Println("Mocking product repository...")
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

type ProductRepositoryMockImplFailed struct {
}

func (pr *ProductRepositoryMockImplFailed) getProductsByCategory(category string) ([]Product, error) {
	fmt.Println("Mocking product repository...")
	return []Product{
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
		productRepository ProductRepository
		expected          int
	}{
		{
			productRepository: &ProductRepositoryMockImplSuccess{},

			expected: 2,
		},
		{
			productRepository: &ProductRepositoryMockImplFailed{},

			expected: 1,
		},
	}
	for _, c := range cases {
		productService := ProductService{
			productRepository: c.productRepository}
		availableProducts, error := productService.getAvaialbleProducts("Electronics")
		if error != nil {
			panic(error)
		}
		if len(availableProducts) != c.expected {
			t.Errorf("Expected %v but got %v", c.expected, len(availableProducts))
		}
	}
}
