package routines

import (
	"fmt"
	"strconv"
	"time"
)

func Service(config ServiceConfig, machines []chan interface{}, broken chan int) {
	var states = make([]bool, len(machines))
	var brokenQueue = make([]int, 0)
	var repairmen = 0
	var repairs = make(chan int)
	for {
		select {
		case brokenMachine := <-broken:
			if !states[brokenMachine] {
				states[brokenMachine] = true
				if repairmen < config.Repairman {
					repairmen++
					go func() {
						time.Sleep(config.Delay)
						if config.Verbose {
							fmt.Println("#[SER " + strconv.Itoa(repairmen-1) + "] fixes " + strconv.Itoa(brokenMachine))
						}
						repairs <- brokenMachine
					}()
				} else {
					if config.Verbose {
						fmt.Println("_[SER H] is adding a machine to broken queue")
					}
					brokenQueue = append(brokenQueue, brokenMachine)
				}
			}
		case repairedMachine := <-repairs:
			repairmen--
			states[repairedMachine] = false
			if len(brokenQueue) > 0 {
				if config.Verbose {
					fmt.Println("_[SER H] is processing the queue")
				}
				repairmen++
				var machine = brokenQueue[0]
				brokenQueue = brokenQueue[1:]
				go func() {
					time.Sleep(config.Delay)
					if config.Verbose {
						fmt.Println("#[SER " + strconv.Itoa(repairmen-1) + "] fixes " + strconv.Itoa(machine))
					}
					repairs <- machine
				}()
			}
		}
	}
}
