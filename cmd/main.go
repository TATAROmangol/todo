package main

import (
	"log/slog"
	"os"
	"todo/internal/config"
)

func main() {
	log := slog.New(
		slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
	)

	cfg := config.MustLoad()
	log.Info("Config loaded", slog.String("cfg", cfg.Env))

	
}
