package repository

type ProductRepositoryStub struct {
}

func (p ProductRepositoryStub) GetName() string {
	return "John Doe"
}
