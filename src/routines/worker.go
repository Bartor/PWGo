package routines

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func Worker(config WorkerConfig, tasks chan chan Task, results chan<- Item, state <-chan interface{}, machines []chan chan Task, service chan int) {
	var tasksDone = 0
	for {
		var req = make(chan Task)
		tasks <- req
		select {
		case task := <-req:
			var machineIdx = rand.Int() % len(machines)

		MachineLoop:
			for {

				var machineChannel = machines[machineIdx]
				var channel = make(chan Task)
				select {
				case machineChannel <- channel:
					channel <- task
					res := <-channel
					if res.Broken {
						fmt.Println("=[WOR " + strconv.Itoa(config.Id) + "] found broken machine " + strconv.Itoa(machineIdx))
						service <- machineIdx
					} else {
						results <- Item{Value: res.Res}
						tasksDone++
						if config.Verbose {
							fmt.Println("=[WOR " + strconv.Itoa(config.Id) + "] did his " + strconv.Itoa(tasksDone) + "# task")
						}
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
