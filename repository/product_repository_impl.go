package repository

import (
	"fmt"
	"go-practice/model"
	"time"
)

type ProductRepositoryImpl struct {
}

func (pr *ProductRepositoryImpl) GetProductsByCategory(category string) ([]model.Product, error) {
	fmt.Println("Connecting to database...")
	time.Sleep(time.Second * 2)

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
