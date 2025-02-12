package router

import (
	"log/slog"
	"net/http"
	"todo/internal/router/handlers"
	"todo/internal/storage"
)

type Router struct{
	log *slog.Logger
	db storage.Storage
	mux *http.ServeMux
}

func New(log *slog.Logger, db storage.Storage) *Router{
	mux := *http.NewServeMux()
	mux.HandleFunc("/create", handlers.Create(log, db))
	mux.HandleFunc("/remove", handlers.Remove(log, db))
	mux.HandleFunc("/getAll", handlers.GetAll(log, db))
	return &Router{log, db, &mux}
}

func (r *Router) Run(path string){
	r.log.Info("Run server", slog.String("path", path))
	http.ListenAndServe(path, r.mux)
}