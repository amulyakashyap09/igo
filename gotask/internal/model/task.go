package model

import (
	"errors"
	"fmt"
	"time"
)

type Status string

const (
	StatusPending   Status = "pending"
	StatusRunning   Status = "running"
	StatusCompleted Status = "completed"
	StatusFailed    Status = "failed"
)

type Task struct {
	ID        int
	Name      string
	Status    Status
	CreatedAt time.Time
}

func NewTask(id int, name string) (*Task, error) {
	if name == "" {
		return nil, errors.New("task name cannot be empty")
	}
	task := &Task{
		ID:        id,
		Name:      name,
		Status:    StatusPending,
		CreatedAt: time.Now(),
	}
	return task, nil
}

func (t *Task) Start() error {
	if t.Status != StatusPending {
		return fmt.Errorf(
			"cannot start task with status %s",
			t.Status,
		)
	}
	t.Status = StatusRunning
	return nil
}

func (t *Task) Complete() error {
	t.Status = StatusCompleted
	return nil
}

func processTask(task *Task) error {
	if err := task.Start(); err != nil {
		return err
	}

	fmt.Printf("processing task: %s\n", task.Name)

	time.Sleep(time.Second)

	task.Complete()

	return nil
}

func Run() {
	tasks := make([]*Task, 0)
	task1, err := NewTask(1, "send-email")
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	task2, err := NewTask(2, "generate-report")
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	tasks = append(tasks, task1, task2)

	for _, task := range tasks {
		if err := processTask(task); err != nil {
			fmt.Printf("error processing task %s: %v\n", task.Name, err)
			continue
		}
		fmt.Printf("task %s completed successfully\n", task.Name)
	}

}
