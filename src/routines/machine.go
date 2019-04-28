package routines

import (
	"fmt"
	"strconv"
	"time"
)

func Machine(config MachineConfig, tasks chan chan Task) {
	for {
		channel := <-tasks
		for task := range channel {
			task.ResolveTask()
			if config.Verbose {
				fmt.Println("~[MCH " + strconv.Itoa(config.Id) + "] finished a task")
			}
			channel <- task
			time.Sleep(config.Delay)
		}
	}
}
