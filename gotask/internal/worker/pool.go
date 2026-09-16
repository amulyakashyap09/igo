package worker

import (
	"fmt"
	"sync"

	"github.com/amulyakashyap09/igo/gotask/internal/service"
)

type Pool struct {
	workerCount int
	jobs        chan int
	service     *service.TaskService
	wg          sync.WaitGroup
}

func NewPool(
	workerCount int,
	queueSize int,
	taskService *service.TaskService,
) *Pool {
	return &Pool{
		workerCount: workerCount,
		jobs:        make(chan int, queueSize),
		service:     taskService,
	}
}

func (p *Pool) Start() {
	for i := 1; i <= p.workerCount; i++ {
		p.wg.Add(1)

		go p.worker(i)
	}
}

func (p *Pool) Submit(taskID int) {
	p.jobs <- taskID
}

func (p *Pool) Stop() {
	close(p.jobs)

	p.wg.Wait()
}

func (p *Pool) worker(workerID int) {
	defer p.wg.Done()

	fmt.Printf("worker %d started\n", workerID)

	for taskID := range p.jobs {
		fmt.Printf(
			"worker %d processing task %d\n",
			workerID,
			taskID,
		)

		err := p.service.ProcessTask(taskID)

		if err != nil {
			fmt.Printf(
				"worker %d failed task %d: %v\n",
				workerID,
				taskID,
				err,
			)
			continue
		}

		fmt.Printf(
			"worker %d completed task %d\n",
			workerID,
			taskID,
		)
	}

	fmt.Printf("worker %d stopped\n", workerID)
}
