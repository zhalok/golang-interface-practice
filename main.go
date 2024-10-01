package main

import (
	"go-practice/repository"
	"go-practice/service"
)

func main() {
	// dependency injection
	productService := service.ProductService{
		ProductRepository: &repository.ProductRepositoryImpl{}}

	availableProducts, error := productService.GetAvaialbleProducts("Electronics")
	if error != nil {
		panic(error)
	}
	for _, product := range availableProducts {
		println(product.Name)
	}
}
