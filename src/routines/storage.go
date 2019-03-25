package routines

import "fmt"

func Tasks(verbose bool, limit int, requests <-chan GetRequestTask, new <-chan Task, state <-chan interface{}) {
	var taskList = make([]Task, 0)
	for {
		select {
		case req := <-requests:
			if len(taskList) != 0 {
				req.Response <- Task{}
			} else {
				req.Response <- taskList[0]
				if verbose {
					fmt.Println("task " + taskList[0].String() + " is given to a worker")
				}
				taskList = taskList[1:]
			}
		case task := <-new:
			if len(taskList) >= limit {
				if verbose {
					fmt.Println("task storage is full")
				}
			} else {
				taskList = append(taskList, task)
			}
		case <-state:
			fmt.Println(taskList)
		}
	}
}

func Items(verbose bool, limit int, requests <-chan GetRequestItem, new <-chan Item, state <-chan interface{}) {
	var itemList = make([]Item, 0)
	for {
		select {
		case req := <-requests:
			if len(itemList) != 0 {
				req.Response <- Item{}
			} else {
				req.Response <- itemList[0]
				if verbose {
					fmt.Println("task " + itemList[0].String() + " is given to a worker")
				}
				itemList = itemList[1:]
			}
		case item := <-new:
			if len(itemList) >= limit {
				if verbose {
					fmt.Println("task storage is full")
				}
			} else {
				itemList = append(itemList, item)
			}
		case <-state:
			fmt.Println(itemList)
		}
	}
}
