package main

import (
	"log/slog"
	"os"
	"todo/internal/config"
	"todo/internal/storage/sqlite"
)

func main() {
	log := slog.New(
		slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
	)

	cfg := config.MustLoad()
	log.Info("Config loaded", slog.String("cfg", cfg.Env))

	db, err := sqlite.New(cfg.StoragePath)
	if err != nil{
		log.Error("Failed in initalize storage",slog.String("err",err.Error()))
	}
	log.Info("Database loaded", slog.String("cfg", cfg.Env))
	_ = db
}
