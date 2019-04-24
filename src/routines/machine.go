package routines

import (
	"fmt"
	"strconv"
	"time"
)

func Machine(config MachineConfig, task chan Task) {
	for {
		t := <-task
		t.ResolveTask()
		time.Sleep(config.Delay)
		if config.Verbose {
			fmt.Println("~[MCH " + strconv.Itoa(config.Id) + "] finished a task")
		}
		task <- t
	}
}
