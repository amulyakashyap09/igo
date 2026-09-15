package service

import (
	"fmt"

	"github.com/amulyakashyap09/igo/gotask/internal/model"
	"github.com/amulyakashyap09/igo/gotask/internal/repository"
)

type TaskService struct {
	repo repository.TaskRepository
}

func NewTaskService(repo repository.TaskRepository) *TaskService {
	return &TaskService{
		repo: repo,
	}
}

func (s *TaskService) CreateTask(id int, name string) (*model.Task, error) {
	task, err := model.NewTask(id, name)

	if err != nil {
		return nil, err
	}

	if err := s.repo.Save(task); err != nil {
		return nil, fmt.Errorf("save task: %w", err)
	}

	return task, nil
}

func (s *TaskService) GetTask(id int) (*model.Task, error) {
	return s.repo.GetByID(id)
}

func (s *TaskService) GetTasks() ([]*model.Task, error) {
	return s.repo.GetAll()
}
