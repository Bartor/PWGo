package routines

import (
	"fmt"
	"time"
)

func Client(id int, verbose bool, offers chan<- chan Item, delay time.Duration) {
	for {
		var req = make(chan Item)
		offers <- req
		var res = <-req

		if res == (Item{}) {
			continue
		} else {
			if verbose {
				fmt.Println("client " + string(id) + " bought " + string(res.Value))
			}
		}

		time.Sleep(delay)
	}
}
