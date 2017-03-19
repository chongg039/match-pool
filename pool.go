package main

import (
	"fmt"
	"sync"
	"time"
)

type Task string

func worker(tasks <-chan Task, quit <-chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case task, ok := <-tasks:
			if !ok {
				return
			}
			fmt.Println("processing task", task)
			time.Sleep(time.Second * 2)
		case <-quit:
			return
		}
	}
}

func main() {
	tasks := make(chan Task, 128)
	quit := make(chan bool)
	var wg sync.WaitGroup

	// spawn 5 workers
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go worker(tasks, quit, &wg)
	}

	// distribute some tasks
	tasks <- Task("foo")
	tasks <- Task("bar")

	// remove two workers
	quit <- true
	quit <- true

	// add three more workers
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go worker(tasks, quit, &wg)
	}

	// distribute more tasks
	for i := 0; i < 20; i++ {
		tasks <- Task(fmt.Sprintf("additional_%d", i+1))
	}

	// end of tasks. the workers should quit afterwards
	close(tasks)
	// use "close(quit)", if you do not want to wait for the remaining tasks

	// wait for all workers to shut down properly
	wg.Wait()
}
