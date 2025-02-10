package main

import (
	"log/slog"
	"os"
	"todo/interfaces"
	"todo/internal/config"
	"todo/internal/storage"
)

func main() {
	cfg := config.MustLoad()

	log := slog.New(
		slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
	)

	log.Info("Config loaded", slog.String("cfg", cfg.Env))

	var db interfaces.Storage
	db, err := sqlite.New(cfg.StoragePath)
	if err != nil{
		log.Error("Failed in initalize storage",slog.String("err",err.Error()))
		os.Exit(1)
	}
	log.Info("Database loaded", slog.String("cfg", cfg.Env))
	_ = db
}
