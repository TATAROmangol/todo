package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
	"todo/internal/config"
	"todo/internal/logger"
	"todo/internal/repository"
	v1 "todo/internal/servers/http/v1"
	"todo/internal/services"
	"todo/pkg/sqlite"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil{
		log.Fatal("failed to load env")
	}

	cfg := config.MustLoad()

	ctx := context.Background()
	ctx = logger.ImportInContext(ctx)

	sqlite, err := sqlite.New(cfg.RepoConfig)
	if err != nil {
		logger.GetFromCtx(ctx).ErrorContext(ctx, "failed in initialize storage", "error", err.Error())
		os.Exit(1)
	}
	logger.GetFromCtx(ctx).Info("database loaded")

	taskRepo := repository.NewRepository(ctx, sqlite)
	taskService := service.NewService(ctx, taskRepo)

	router := v1.New(ctx, cfg.HttpConfig, taskService)

	go func(){
		if err := router.Run(); err != nil{
			logger.GetFromCtx(ctx).ErrorContext(ctx, "failed in server", "error", err.Error())
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
}
