package routines

import (
	"fmt"
	"strconv"
	"time"
)

func Machine(config MachineConfig, tasks chan chan Task) {
	for {
		channel := <-tasks
		select {
		case task := <-channel:
			task.ResolveTask()
			time.Sleep(config.Delay)
			if config.Verbose {
				fmt.Println("~[MCH " + strconv.Itoa(config.Id) + "] finished a task")
			}
			channel <- task
		default:
			fmt.Println("xD")
		}
	}
}
