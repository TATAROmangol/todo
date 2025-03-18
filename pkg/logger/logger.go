package logger

import (
	"context"
	"log/slog"
	"os"
)

const(
	Key = "logger"
)

func ImportInContext(ctx context.Context) context.Context{
	log := slog.New(
		slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
	)

	return context.WithValue(ctx, Key, log)
}

func GetFromCtx(ctx context.Context) *slog.Logger{
	return ctx.Value(Key).(*slog.Logger)
}
