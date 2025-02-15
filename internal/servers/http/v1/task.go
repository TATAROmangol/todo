package v1

import (
	"encoding/json"
	"net/http"
	"todo/internal/entities"
)

type TaskHandler struct{
	service Service
}

func NewTaskHandler(service Service) *TaskHandler{
	return &TaskHandler{service}
}

type CreateReq struct{
	Name string `json:"name"`
}

func (th *TaskHandler) Create(w http.ResponseWriter, r *http.Request){
	var req CreateReq
	err := json.NewDecoder(r.Body).Decode(&req)
	defer r.Body.Close()
	if err != nil {
		WriteError(w, err, 404)
		return
	}

	task, err := th.service.CreateTask(req.Name)
	if err != nil {
		WriteError(w, err, 404)
		return
	}

	err = json.NewEncoder(w).Encode(task)
	if err != nil {
		WriteError(w, err, 404)
		return
	}
}

func (th *TaskHandler) Remove(w http.ResponseWriter, r *http.Request){
	var task entities.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
			WriteError(w, err, 404)
			return
		}
		defer r.Body.Close()

		if err := th.service.RemoveTask(task.Id); err != nil {
			WriteError(w, err, 404)
			return
		}
}

func (th *TaskHandler) Get(w http.ResponseWriter, r *http.Request){
	tasks, err := th.service.GetTasks()
		if err != nil {
			WriteError(w, err, 404)
			return
		}

		err = json.NewEncoder(w).Encode(tasks)
		if err != nil {
			WriteError(w, err, 404)
			return
		}
}
