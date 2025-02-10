package interfaces

import "todo/models"

type Storage interface {
	GetTasks() ([]models.Task, error)
	AddTask(models.Task) error
	RemoveTask(int) error
}