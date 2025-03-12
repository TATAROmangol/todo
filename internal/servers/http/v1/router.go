package v1

import (
	"context"
	"net/http"
	"todo/internal/logger"
)

type Service interface {
	TaskService
}

type Router struct {
	ctx context.Context
	cfg Config
	srv *http.Server
}

func New(ctx context.Context, cfg Config, cases Service) *Router {
	th := NewTaskHandler(ctx, cases)

	mux := http.NewServeMux()
	mux.HandleFunc("/", th.Hello)
	mux.HandleFunc("/create", th.Create)
	mux.HandleFunc("/remove", th.Remove)
	mux.HandleFunc("/get", th.Get)

	srv := &http.Server{
		Addr: cfg.Address,
		Handler: mux,
	}
	return &Router{ctx, cfg, srv}
}

func (r *Router) Run() error{
	logger.GetFromCtx(r.ctx).Info("Run server", "path", r.cfg.Address)
	if err := r.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed{
		return err
	}
	return nil
}

func (r *Router) Shutdown(ctx context.Context) error{
	if err := r.srv.Shutdown(ctx); err != nil{
		return err
	}
	return nil
}
