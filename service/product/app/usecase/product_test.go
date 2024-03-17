package usecase

import (
	"product/app/repository"
	"testing"
)

func TestProduct(t *testing.T) {
	tests := []struct {
		name     string
		expected string
	}{
		{"should return `John Doe` value", "John Doe"},
	}

	productStub := NewProduct(&repository.ProductRepositoryStub{})

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			expected := productStub.repository.GetName()
			if expected != "John Doe" {
				t.Error("error nich")
			}
		})
	}
}
