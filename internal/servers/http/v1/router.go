package v1

import (
	"log/slog"
	"net/http"
	"todo/internal/entities"
)

type Service interface{
	GetTasks() ([]entities.Task, error)
	CreateTask(name string) (entities.Task, error)
	RemoveTask(id int) error
}

type Router struct{
	log *slog.Logger
	mux *http.ServeMux
}

func New(log *slog.Logger, service Service) *Router{
	th := NewTaskHandler(service)
	mux := *http.NewServeMux()
	mux.HandleFunc("/create", th.Create)
	mux.HandleFunc("/remove", th.Remove)
	mux.HandleFunc("/get", th.Get)
	return &Router{log, &mux}
}

func (r *Router) Run(path string){
	r.log.Info("Run server", slog.String("path", path))
	http.ListenAndServe(path, r.mux)
}