package gotask

import (
	"fmt"

	"github.com/amulyakashyap09/igo/gotask/internal/repository"
	"github.com/amulyakashyap09/igo/gotask/internal/service"
	"github.com/amulyakashyap09/igo/gotask/internal/worker"
)

func Run() {

	repo := repository.NewMemoryTaskRepository()
	taskService := service.NewTaskService(repo)

	// task1, err := service.CreateTask(1, "send-email")
	// if err != nil {
	// 	fmt.Println("error:", err)
	// 	return
	// }

	// task2, err := service.CreateTask(2, "generate-report")
	// if err != nil {
	// 	fmt.Println("error:", err)
	// 	return
	// }

	// fmt.Println("created:", task1)
	// fmt.Println("created:", task2)

	// tasks, err := service.GetTasks()
	// if err != nil {
	// 	fmt.Println("error:", err)
	// 	return
	// }

	// for _, task := range tasks {
	// 	fmt.Printf("task %d: %s (status: %s)\n", task.ID, task.Name, task.Status)
	// }

	for i := 1; i <= 10; i++ {
		_, err := taskService.CreateTask(
			i,
			fmt.Sprintf("task-%d", i),
		)

		if err != nil {
			fmt.Println("create task:", err)
			return
		}
	}

	pool := worker.NewPool(
		3,
		10,
		taskService,
	)

	pool.Start()

	for i := 1; i <= 10; i++ {
		pool.Submit(i)
	}

	pool.Stop()

	fmt.Println("all tasks finished")

}
