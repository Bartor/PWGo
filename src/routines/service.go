package routines

import (
	"fmt"
	"strconv"
	"time"
)

func Service(config ServiceConfig, machines []chan interface{}, broken chan int) {
	var states = make([]bool, len(machines))
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
				}
			}
		case repairedMachine := <-repairs:
			repairmen--
			states[repairedMachine] = false
		}
	}
}
