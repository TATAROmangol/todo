package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"todo/internal/entities"
	"todo/internal/storage"
)

func Create(log *slog.Logger, db storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		var temp struct{Name string `json:"name"`}
		err := json.NewDecoder(r.Body).Decode(&temp)
		defer r.Body.Close()
		if err != nil{
			WriteError(w, err, 404)
			return
		}

		task, err := db.CreateTask(temp.Name)
		if err != nil{
			WriteError(w, err, 404)
			return
		}

		err = json.NewEncoder(w).Encode(task)
		if err != nil{
			WriteError(w, err, 404)
			return
		}
	}	
}

func Remove(log *slog.Logger, db storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		var task entities.Task
		err := json.NewDecoder(r.Body).Decode(&task)
		defer r.Body.Close()
		if err != nil{
			WriteError(w, err, 404)
			return
		}

		err = db.RemoveTask(task.Id)
		if err != nil{
			WriteError(w, err, 404)
			return
		}
	}	
}

func GetAll(log *slog.Logger, db storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		tasks, err := db.GetTasks()
		if err != nil{
			WriteError(w, err, 404)
			return
		}

		err = json.NewEncoder(w).Encode(tasks)
		if err != nil{
			WriteError(w, err, 404)
			return
		}
	}	
}

