package routines

import (
	"PWGo/src/conf"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func Machine(config MachineConfig, tasks chan chan Task, backdoor chan interface{}) {
	var working = true
	for {
		select {
		case channel := <-tasks:
			for task := range channel {
				if working {
					task.Broken = false
					task.ResolveTask()
				} else {
					task.Broken = true
				}
				if config.Verbose {
					fmt.Println("~[MCH " + strconv.Itoa(config.Id) + "] finished a task")
				}
				channel <- task
				working = rand.Float32() > conf.BreakingProb
				time.Sleep(config.Delay)
			}
		case <-backdoor:
			working = true
		}
	}
}
