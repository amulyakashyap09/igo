package gotask

import (
	"fmt"

	"github.com/amulyakashyap09/igo/gotask/internal/repository"
	"github.com/amulyakashyap09/igo/gotask/internal/service"
)

func Run() {

	repo := repository.NewMemoryTaskRepository()
	service := service.NewTaskService(repo)

	task1, err := service.CreateTask(1, "send-email")
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	task2, err := service.CreateTask(2, "generate-report")
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println("created:", task1)
	fmt.Println("created:", task2)

	tasks, err := service.GetTasks()
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	for _, task := range tasks {
		fmt.Printf("task %d: %s (status: %s)\n", task.ID, task.Name, task.Status)
	}

}
