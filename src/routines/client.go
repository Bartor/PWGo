package routines

import (
	"fmt"
	"time"
)

func Client(id int, verbose bool, offers chan<- GetRequestItem, delay time.Duration) {
	for {
		var req = GetRequestItem{Response: make(chan Item)}
		offers <- req
		var res = <-req.Response

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
