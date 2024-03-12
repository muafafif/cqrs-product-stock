package usecase

import (
	"context"
	"product/app/repository"
	"product/model"
)

func CreateProduct(ctx context.Context, payload model.CreateProductPayload) error {
	return repository.InsertProduct(context.Background(), model.Product{})
}
