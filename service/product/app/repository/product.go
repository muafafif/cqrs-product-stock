package repository

import "log"

type ProductRepository struct {
}

func NewProduct() ProductRepository {
	return ProductRepository{}
}

func (p ProductRepository) Test() {
	log.Println("Test Repository")
}
