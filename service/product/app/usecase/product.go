package usecase

import "log"

type IProductRepository interface {
	Test()
}

type ProductUsecase struct {
	repository IProductRepository
}

func NewProduct(repository IProductRepository) ProductUsecase {
	return ProductUsecase{repository: repository}
}

func (p ProductUsecase) Test() {
	log.Println("Test Usecase")
	p.repository.Test()
}
