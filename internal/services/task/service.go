package task

import "todo/internal/entities"

type Repo interface {
	GetTasks() ([]entities.Task, error)
	CreateTask(name string) (entities.Task, error)
	RemoveTask(id int) error
}

type Service struct {
	db Repo
}

func NewService(db Repo) *Service {
	return &Service{db}
}

func (tc *Service) GetTasks() ([]entities.Task, error) {
	return tc.db.GetTasks()
}

func (tc *Service) CreateTask(name string) (entities.Task, error) {
	return tc.db.CreateTask(name)
}

func (tc *Service) RemoveTask(id int) error {
	return tc.db.RemoveTask(id)
}
