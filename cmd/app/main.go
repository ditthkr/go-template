package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"go-template/internal/app"
	"go.uber.org/fx"
)

func main() {
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	fxApp := fx.New(app.Module)

	go func() {
		<-shutdown
		log.Println("🛑 Gracefully shutting down...")
		fxApp.Stop(context.Background())
	}()

	fxApp.Run()
}
