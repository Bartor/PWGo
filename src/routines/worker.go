package routines

import (
	"fmt"
	"time"
)

func Worker(id int, verbose bool, tasks chan<- chan Task, results chan<- Item, delay time.Duration) {
	for {
		var req = make(chan Task)
		tasks <- req
		var res = <-req

		if result, e := res.ResolveTask(); e == nil {
			continue
		} else {
			results <- Item{Value: result}
			if verbose {
				fmt.Println("worker " + string(id) + " resolved a task and sends " + string(result))
			}
		}

		time.Sleep(delay)
	}
}
