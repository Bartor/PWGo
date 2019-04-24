package routines

import (
	"fmt"
	"strconv"
	"time"
)

func Worker(config WorkerConfig, tasks chan chan Task, results chan<- Item, state <-chan interface{}, machines []chan Task) {
	var tasksDone = 0
	for {
		var req = make(chan Task)
		tasks <- req
		select {
		case task := <-req:
			var machineIdx = 0

		MachineLoop:
			for {
				var machineChannel = machines[machineIdx]

				select {
				case machineChannel <- task:
					res := <-machineChannel
					results <- Item{Value: res.Res}
					tasksDone = (tasksDone + 1) % len(machines)
					if config.Verbose {
						fmt.Println("=[WOR " + strconv.Itoa(config.Id) + "] : " + strconv.Itoa(tasksDone))
					}
					break MachineLoop
				case <-time.After(config.Timeout):
					machineIdx++
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
