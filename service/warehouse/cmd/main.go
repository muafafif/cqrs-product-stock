package main

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	logger := slog.Default()
	logger.Info("STARTING THE SERVICE")

	//appContext := context.Background()

	sig := make(chan os.Signal, 0)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig

	logger.Warn("STOPING THE SERVICE")
}
