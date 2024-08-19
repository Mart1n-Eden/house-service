package main

import (
	"context"
	"flag"
	"log/slog"
	"os"
	"os/signal"
	"time"

	"house-service/internal/cache"
	"house-service/internal/config"
	"house-service/internal/http/handler"
	"house-service/internal/http/server"
	"house-service/internal/logger"
	"house-service/internal/repository"
	"house-service/internal/sender"
	"house-service/internal/service/auth"
	"house-service/internal/service/flat"
	"house-service/internal/service/house"
	"house-service/internal/service/subscribe"
	"house-service/internal/token"
)

const (
	shutdownTimeout = 5 * time.Second
)

func main() {
	configPath := flag.String("config", "config/config.yaml", "config file path")
	flag.Parse()

	cfg := config.ParseConfig(*configPath)

	logger.MustInit(cfg.Logger.Level)

	pg, err := repository.NewConnection(context.Background(), cfg.DB)
	if err != nil {
		logger.Error("error connecting to database", slog.String("error", err.Error()))

		os.Exit(1)
	}

	repo := repository.New(pg)
	c := cache.New()

	tok := token.New(cfg.Secret)

	send := sender.New()

	houseService := house.New(repo)
	flatService := flat.New(repo, c)
	authService := auth.New(repo, tok)
	subService := subscribe.New(repo, send)

	go func() {
		ticker := time.NewTicker(12 * time.Hour)
		defer ticker.Stop()

		for range ticker.C {
			go subService.GetHouseBySubscription(context.Background())
		}
	}()

	hnd := handler.New(houseService, flatService, authService, subService)

	app := server.New(hnd.Route(), cfg.Server)

	go func() {
		if err = app.Run(); err != nil {
			logger.Error("error running server", slog.String("error", err.Error()))
		}
	}()

	logger.Info("starting server", slog.String("address", cfg.Server.Address))

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)
	<-quit

	logger.Info("stopping server")

	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	if err = app.Shutdown(ctx); err != nil {
		logger.Error("shutdown server error", slog.String("error", err.Error()))
	}
}
