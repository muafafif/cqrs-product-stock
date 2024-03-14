package apihandler

import (
	"net/http"
	"product/app/usecase"
	"product/model"
)

func SetAPIHandler(mux *http.ServeMux) {
	// SET ROUTES
	mux.HandleFunc("POST /product", createProduct)
}

func createProduct(w http.ResponseWriter, r *http.Request) {
	usecase.CreateProduct(r.Context(), model.CreateProductPayload{})
}
