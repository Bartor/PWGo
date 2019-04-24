package main

import (
	"PWGo/src/conf"
	"PWGo/src/routines"
	"fmt"
	"math/rand"
)

func main() {
	var taskListState = make(chan interface{})
	var itemListState = make(chan interface{})
	var workersStates = make([]chan interface{}, 0)

	for i := 0; i < conf.Workers; i++ {
		var channel = make(chan interface{})
		workersStates = append(workersStates, channel)
	}

	var ceoToTaskList = make(chan routines.Task)
	var taskListToWorkers = make(chan chan routines.Task)
	var workersToItemList = make(chan routines.Item)
	var itemListToClients = make(chan chan routines.Item)

	var machines = make([]chan chan routines.Task, 0)
	for i := 0; i < conf.Machines; i++ {
		var machineChan = make(chan chan routines.Task)
		machines = append(machines, machineChan)
		go routines.Machine(routines.MachineConfig{i, conf.Verbose, conf.DelayMachine}, machineChan)
	}

	go routines.Tasks(conf.TaskListSize, ceoToTaskList, taskListToWorkers, taskListState)
	go routines.Items(conf.ItemListSize, workersToItemList, itemListToClients, itemListState)

	for i := 0; i < conf.Workers; i++ {
		go routines.Worker(routines.WorkerConfig{
			Id:      i,
			Verbose: conf.Verbose,
			Delay:   conf.DelayWorker,
			Timeout: conf.TimeoutWorker,
			Patient: rand.Int()%2 == 0}, taskListToWorkers, workersToItemList, workersStates[i], machines)
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
