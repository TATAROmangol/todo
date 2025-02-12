package main

import (
	"log/slog"
	"os"
	"todo/internal/config"
	"todo/internal/router"
	"todo/internal/storage"
	"todo/internal/storage/sqlite"
)

func main() {
	cfg := config.MustLoad()

	log := slog.New(
		slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
	)

	log.Info("Config loaded", slog.String("cfg", cfg.Env))

	var db storage.Storage
	db, err := sqlite.New(cfg.StoragePath)
	if err != nil{
		log.Error("Failed in initalize storage",slog.String("err",err.Error()))
		os.Exit(1)
	}
	
	log.Info("Database loaded", slog.String("cfg", cfg.Env))
	
	router := router.New(log, db)
	router.Run(cfg.Address)
}
