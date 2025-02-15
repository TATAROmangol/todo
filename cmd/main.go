package main

import (
	"log/slog"
	"os"
	"todo/internal/config"
	v1 "todo/internal/servers/http/v1"
	"todo/internal/services/task"
	"todo/pkg/sqlite"
)

func main() {
	cfg := config.MustLoad()

	log := slog.New(
		slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
	)

	log.Info("Config loaded", slog.String("cfg", cfg.Env))

	sqlite, err := sqlite.New(cfg.StoragePath)
	if err != nil {
		log.Error("Failed in initalize storage", slog.String("err", err.Error()))
		os.Exit(1)
	}

	log.Info("Database loaded", slog.String("cfg", cfg.Env))

	taskRepo := task.NewRepository(sqlite)
	taskService := task.NewService(taskRepo)

	router := v1.New(log, taskService)
	router.Run(cfg.Address)
}
