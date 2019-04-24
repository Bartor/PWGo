package routines

import (
	"fmt"
	"strconv"
	"time"
)

//todo add machines
func Worker(config WorkerConfig, tasks chan chan Task, results chan<- Item, state <-chan interface{}) {
	var tasksDone = 0
	for {
		var req = make(chan Task)
		tasks <- req
		select {
		case res := <-req:
			if result, e := res.ResolveTask(); e != nil {
				//this shouldn't happen now
				continue
			} else {
				results <- Item{Value: result}
				tasksDone++
				if config.Verbose {
					fmt.Println("=[WOR " + strconv.Itoa(config.Id) + "] : " + strconv.Itoa(tasksDone))
				}
			}

			close(req)
			time.Sleep(config.Delay)
		case <-state:
			if config.Patient {
				fmt.Println("=[WOR " + strconv.Itoa(config.Id) + "] is patient and did " + strconv.Itoa(tasksDone))
			} else {
				fmt.Println("=[WOR " + strconv.Itoa(config.Id) + "] is impatient and did " + strconv.Itoa(tasksDone))
			}
		}
	}
}
