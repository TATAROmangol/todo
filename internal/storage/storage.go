package storage

import "todo/internal/models"

type Storage interface {
	New(string) (*Storage, error)
	GetTasks() []models.Task
	AddTask(models.Task) error
	RemoveTask(int) error
}




