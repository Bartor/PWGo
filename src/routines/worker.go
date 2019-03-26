package routines

import (
	"fmt"
	"strconv"
	"time"
)

func Worker(id int, verbose bool, tasks chan chan Task, results chan<- Item, delay time.Duration) {
	for {
		var req = make(chan Task)
		tasks <- req
		var res = <-req

		if result, e := res.ResolveTask(); e != nil {
			//this shouldn't happen now
			continue
		} else {
			results <- Item{Value: result}
			if verbose {
				fmt.Println("=[WOR " + strconv.Itoa(id) + "] resolved a task and sends " + strconv.Itoa(result))
			}
		}

		time.Sleep(delay)
	}
}
