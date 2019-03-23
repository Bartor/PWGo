package structs

import "time"

type Worker struct {
	Delay   int
	Tasks   chan Task
	Results chan int
}

func (worker *Worker) Start() {
	for {
		time.Sleep(time.Duration(worker.Delay))
		t := <-worker.Tasks
		if r, e := t.ResolveTask(); e != nil {
			worker.Results <- r
		}
	}
}
