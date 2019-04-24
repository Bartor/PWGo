package routines

import (
	"fmt"
	"strconv"
	"time"
)

func Machine(config MachineConfig, tasks chan Task) {
	for {
		t := <-tasks
		t.ResolveTask()
		time.Sleep(config.Delay)
		if config.Verbose {
			fmt.Println("~[MCH " + strconv.Itoa(config.Id) + "] finished a task")
		}
		tasks <- t
	}
}
