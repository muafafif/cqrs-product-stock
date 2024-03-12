package server

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"product/presentation/api"
	"syscall"
)

func Start() {
	logger := slog.Default()
	mux := http.NewServeMux()

	api.Set(mux)

	server := http.Server{
		Handler: mux,
		Addr:    ":8080",
	}

	go func() {
		sig := make(chan os.Signal, 0)
		signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
		<-sig

		logger.Warn("STOPING THE SERVICE")
		server.Shutdown(context.Background())
	}()

	logger.Info("STARTING THE SERVICE")
	server.ListenAndServe()
}
