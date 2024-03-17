package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	apihandler "product/app/handler/api"
	"product/app/repository"
	"product/app/usecase"
	"syscall"
)

func main() {
	log.Println("STARTING")
	sig := make(chan os.Signal, 1)

	productRepository := repository.NewProduct()
	productUsecase := usecase.NewProduct(&productRepository)
	ProductAPIHandler := apihandler.NewProduct(&productUsecase)

	mux := http.NewServeMux()

	ProductAPIHandler.Router(mux)

	server := http.Server{
		Handler: mux,
		Addr:    ":8080",
	}

	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)
	<-sig
	log.Println("STOPPING")

	server.ListenAndServe()
	server.Shutdown(context.Background())
}
