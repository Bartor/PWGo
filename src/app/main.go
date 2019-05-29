package main

import (
	"PWGo/src/conf"
	"PWGo/src/routines"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
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

	var service = make(chan int)

	var machines = make([]chan chan routines.Task, 0)
	var backdoors = make([]chan interface{}, 0)

	for i := 0; i < conf.Machines; i++ {
		var machineChan = make(chan chan routines.Task)
		var backdoor = make(chan interface{})
		machines = append(machines, machineChan)
		backdoors = append(backdoors, backdoor)
		go routines.Machine(routines.MachineConfig{Id: i, Verbose: conf.Verbose, Delay: conf.DelayMachine, Prob: conf.BreakingProb}, machineChan, backdoor)
	}

	go routines.Tasks(conf.TaskListSize, ceoToTaskList, taskListToWorkers, taskListState)
	go routines.Items(conf.ItemListSize, workersToItemList, itemListToClients, itemListState)
	go routines.Service(routines.ServiceConfig{Repairman: conf.Repairman, Verbose: conf.Verbose, Delay: conf.RepairTime}, backdoors, service)

	for i := 0; i < conf.Workers; i++ {
		go routines.Worker(routines.WorkerConfig{
			Id:      i,
			Verbose: conf.Verbose,
			Delay:   conf.DelayWorker,
			Timeout: conf.TimeoutWorker,
			Patient: rand.Int()%2 == 0}, taskListToWorkers, workersToItemList, workersStates[i], machines, service)
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
			var first = strings.Split(line, ":")
			switch first[0] {
			case "t":
				taskListState <- true
			case "i":
				itemListState <- true
			case "s":
				num, err := strconv.Atoi(first[1])
				if err == nil {
					workersStates[num%len(workersStates)] <- true
				}
			case "h":
				fmt.Println("t - tasks list\ni - item list")
			default:
				fmt.Println("for help, type h")
			}
		}
	}
}
