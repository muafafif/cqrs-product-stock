package apihandler

import (
	"log"
	"net/http"
)

type IProductUsecase interface {
}

type ProductAPIHandler struct {
	usecase IProductUsecase
}

func (p *ProductAPIHandler) Router(mux *http.ServeMux) {
	mux.HandleFunc("GET /product/test", p.Test)
}

func NewProduct(usecase IProductUsecase) ProductAPIHandler {
	return ProductAPIHandler{usecase: usecase}
}

func (p ProductAPIHandler) Test(w http.ResponseWriter, r *http.Request) {
	log.Println("Test API Handler")
	w.Write([]byte("Success Test"))
}
