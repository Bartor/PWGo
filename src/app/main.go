package main

import (
	"PWGo/src/conf"
	"PWGo/src/routines"
	"fmt"
)

func main() {
	var taskListState = make(chan interface{})
	var itemListState = make(chan interface{})

	var ceoToTaskList = make(chan routines.Task)
	var taskListToWorkers = make(chan chan routines.Task)
	var workersToItemList = make(chan routines.Item)
	var itemListToClients = make(chan chan routines.Item)

	go routines.Tasks(conf.Verbose, conf.TaskListSize, taskListToWorkers, ceoToTaskList, taskListState)
	go routines.Items(conf.Verbose, conf.ItemListSize, itemListToClients, workersToItemList, itemListState)

	for i := 0; i < conf.Workers; i++ {
		go routines.Worker(i, conf.Verbose, taskListToWorkers, workersToItemList, conf.DelayWorker)
	}

	for i := 0; i < conf.Clients; i++ {
		go routines.Client(i, conf.Verbose, itemListToClients, conf.DelayClient)
	}

	go routines.Ceo(conf.Verbose, ceoToTaskList, conf.DelayCeoLo, conf.DelayCeoHi)

	if conf.Verbose {
		fmt.Scanln()
	} else {
		for {
			var line string
			fmt.Scanln(&line)
			switch line {
			case "t":
				taskListState <- true
			case "i":
				itemListState <- true
			case "h":
				fmt.Println("t - tasks list\ni - item list")
			default:
				fmt.Println("for help, type h")
			}
		}
	}
}
