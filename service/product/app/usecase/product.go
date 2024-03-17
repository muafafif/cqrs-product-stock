package usecase

type IProductRepository interface {
	GetName() string
}

type ProductUsecase struct {
	repository IProductRepository
}

func NewProduct(repository IProductRepository) ProductUsecase {
	return ProductUsecase{repository: repository}
}

// func (p ProductUsecase) Test() {
// 	log.Println("Test Usecase")
// 	p.repository.Test()
// }

func (p ProductUsecase) VerifyName() string {
	return p.repository.GetName()
}
