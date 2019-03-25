package routines

import (
	"fmt"
	"strconv"
	"time"
)

func Client(id int, verbose bool, offers chan chan Item, delay time.Duration) {
	for {
		var req = make(chan Item)
		offers <- req
		var res = <-req

		if res == (Item{}) {
			continue
		} else {
			if verbose {
				fmt.Println("==[CLI " + strconv.Itoa(id) + "] bought " + strconv.Itoa(res.Value))
			}
		}

		time.Sleep(delay)
	}
}
