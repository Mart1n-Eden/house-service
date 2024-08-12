package main

import (
	"context"
	"fmt"
	"log"

	"house-service/internal/config"
	"house-service/internal/repository"
)

func main() {
	cfg := config.ParseConfig("config/config.yaml")

	pg, err := repository.NewConnection(context.Background(), cfg.DB)
	if err != nil {
		log.Fatal("error connecting to database", err)
	}

	repo := repository.New(pg)

	fmt.Println(repo) // pass
}
