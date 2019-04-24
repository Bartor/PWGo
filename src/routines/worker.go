package routines

import (
	"fmt"
	"strconv"
	"time"
)

func Worker(config WorkerConfig, tasks chan chan Task, results chan<- Item, state <-chan interface{}, machines []chan chan Task) {
	var tasksDone = 0
	for {
		var req = make(chan Task)
		tasks <- req
		select {
		case task := <-req:
			var machineIdx = config.Id % len(machines)

		MachineLoop:
			for {
				var machineChannel = machines[machineIdx]
				var channel = make(chan Task)

				select {
				case machineChannel <- channel:
					channel <- task
					res := <-channel
					results <- Item{Value: res.Res}
					tasksDone++
					if config.Verbose {
						fmt.Println("=[WOR " + strconv.Itoa(config.Id) + "] did his " + strconv.Itoa(tasksDone) + "# task")
					}
					close(channel)
					break MachineLoop
				case <-time.After(config.Timeout):
					if config.Verbose {
						fmt.Println("=[WOR " + strconv.Itoa(config.Id) + "] tries next machine")
					}
					machineIdx = (machineIdx + 1) % len(machines)
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
