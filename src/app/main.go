package main

import (
	"PWGo/src/structs"
	"time"
)

func main() {
	tasksQueue := structs.ThreadListTask{List: make([]structs.Task, 1), In: make(chan<- structs.Task), Out: make(<-chan structs.Task)}
	solveQueue := structs.ThreadListInt{List: make([]int, 1), In: make(chan<- int), Out: make(<-chan int)}

	go tasksQueue.Start()
	go solveQueue.Start()

	ceo := structs.Ceo{HiLimit: 10, LoLimit: 1, Tasks: tasksQueue.In}
	go ceo.Start()

	for i := 0; i < 10; i++ {
		worker := structs.Worker{Delay: 1, Tasks: tasksQueue.Out, Results: solveQueue.In}
		go worker.Start()
	}

	for {
		time.Sleep(time.Second)
	}
}
