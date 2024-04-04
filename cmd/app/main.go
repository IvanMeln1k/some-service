package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/IvanMeln1k/some-service/internal/handler"
	"github.com/IvanMeln1k/some-service/internal/server"
	"github.com/sirupsen/logrus"
)

func main() {
	handlers := handler.NewHandler(handler.Deps{})
	srv := new(server.Server)

	go func() {
		if err := srv.Run(server.ServerConfig{Host: "localhost", Port: "8000"},
			handlers.InitRoutes()); err != nil {
			if !errors.Is(http.ErrServerClosed, err) {
				logrus.Fatalf("error running server: %s", err)
			}
		}
	}()
	logrus.Print("Server starting...")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("Server shutting down...")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Fatalf("error server shutting down: %s", err)
	}

	logrus.Print("Server stoped")
}
