package repository

import "go-practice/model"

type ProductRepository interface {
	GetProductsByCategory(category string) ([]model.Product, error)
}
