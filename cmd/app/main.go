package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"time"

	"house-service/internal/config"
	"house-service/internal/http/handler"
	"house-service/internal/http/server"
	"house-service/internal/logger"
	"house-service/internal/repository"
	"house-service/internal/service/auth"
	"house-service/internal/service/flat"
	"house-service/internal/service/house"
)

const (
	shutdownTimeout = 5 * time.Second
)

func main() {
	cfg := config.ParseConfig("config/config.yaml")

	log := logger.New(cfg.Logger.Level)

	// TODO: change context
	pg, err := repository.NewConnection(context.Background(), cfg.DB)
	if err != nil {
		log.Error("error connecting to database", slog.String("error", err.Error()))

		os.Exit(1)
	}

	repo := repository.New(pg)
	houseSrvc := house.New(repo)
	flatSrvc := flat.New(repo)
	authSrvc := auth.New(repo, cfg.Secret)
	hnd := handler.New(log, houseSrvc, flatSrvc, authSrvc)

	app := server.New(hnd.Route(cfg.Secret), cfg.Server)

	go func() {
		if err := app.Run(); err != nil {
			log.Error("error running server", slog.String("error", err.Error()))
		}
	}()

	log.Info("starting server", slog.String("address", cfg.Server.Address))

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)
	<-quit

	log.Info("stopping server")

	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	if err := app.Shutdown(ctx); err != nil {
		log.Error("shutdown server error", slog.String("error", err.Error()))
	}
}
