package main

import (
	"context"
	"log"
	"mediatrack/config"
	"mediatrack/internal/application"
)

func main() {
	ctx := context.Background()

	cfg, err := config.New()
	if err != nil {
		log.Fatal("failed to init config err:", err)
	}

	app := application.New(ctx, cfg)

	app.InitLogger()

	app.Run()
}
