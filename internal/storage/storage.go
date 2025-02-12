package storage

import "todo/internal/entities"

type Storage interface {
	GetTasks() ([]entities.Task, error)
	CreateTask(string) (entities.Task, error)
	RemoveTask(int) error
}
