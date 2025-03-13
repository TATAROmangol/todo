package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"
	"todo/internal/config"
	"todo/internal/logger"
	"todo/internal/migrator"
	"todo/internal/repository"
	v1 "todo/internal/servers/http/v1"
	"todo/internal/services"
	"todo/pkg/postgres"

	// "github.com/joho/godotenv"
)

func main() {
	// godotenv.Load()
	cfg := config.MustLoad()

	ctx := context.Background()
	ctx = logger.ImportInContext(ctx)

	pq, err := postgres.NewDB(cfg.RepoConfig)
	if err != nil {
		logger.GetFromCtx(ctx).ErrorContext(ctx, "failed in initialize storage", "error", err.Error())
		os.Exit(1)
	}
	logger.GetFromCtx(ctx).Info("database loaded")

	m, err := migrator.New(cfg.RepoConfig)
	if err != nil{
		logger.GetFromCtx(ctx).ErrorContext(ctx, "failed in create migrator", "error", err.Error())
		os.Exit(1)
	}

	if err := m.Up(); err != nil{
		logger.GetFromCtx(ctx).ErrorContext(ctx, "failed in up migrate", "error", err.Error())
		os.Exit(1)
	}

	taskRepo := repository.NewRepository(ctx, pq)
	taskService := service.NewService(ctx, taskRepo)

	router := v1.New(ctx, cfg.HttpConfig, taskService)

	go func(){
		if err := router.Run(); err != nil{
			logger.GetFromCtx(ctx).ErrorContext(ctx, "failed in server", "error", err.Error())
			os.Exit(1)
		}
	}()

	c := make(chan os.Signal, 1)
    signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)

	<-c

	logger.GetFromCtx(ctx).Info("started shutdown")

	closeCtx, cancel := context.WithTimeout(ctx, 10 * time.Second)
	defer cancel()
	
	router.Shutdown(closeCtx)
	logger.GetFromCtx(ctx).Info("server stop")

	pq.Close()
}
