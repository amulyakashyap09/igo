package repository

import (
	"errors"
	"sync"

	"github.com/amulyakashyap09/igo/gotask/internal/model"
)

var ErrTaskNotFound = errors.New("task not found")

type TaskRepository interface {
	Save(task *model.Task) error
	GetByID(id int) (*model.Task, error)
	GetAll() ([]*model.Task, error)
}

type MemoryTaskRepository struct {
	mu    sync.RWMutex
	tasks map[int]*model.Task
}

func NewMemoryTaskRepository() *MemoryTaskRepository {
	return &MemoryTaskRepository{
		tasks: make(map[int]*model.Task),
	}
}

func (r *MemoryTaskRepository) Save(task *model.Task) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.tasks[task.ID] = task

	return nil
}

func (r *MemoryTaskRepository) GetByID(id int) (*model.Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	task, exists := r.tasks[id]

	if !exists {
		return nil, ErrTaskNotFound
	}

	return task, nil
}

func (r *MemoryTaskRepository) GetAll() ([]*model.Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	tasks := make([]*model.Task, 0, len(r.tasks))

	for _, task := range r.tasks {
		tasks = append(tasks, task)
	}

	return tasks, nil
}
