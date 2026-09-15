package service

import (
	"fmt"
	"time"

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

func (s *TaskService) ProcessTask(id int) error {
	task, err := s.repo.GetByID(id)

	if err != nil {
		return fmt.Errorf("get task: %w", err)
	}

	if err := task.Start(); err != nil {
		return fmt.Errorf("start task: %w", err)
	}

	fmt.Printf("processing task %d: %s\n", task.ID, task.Name)

	time.Sleep(2 * time.Second)

	task.Complete()

	if err := s.repo.Save(task); err != nil {
		return fmt.Errorf("save completed task: %w", err)
	}

	return nil
}
