package routines

import (
	"fmt"
	"sync"
)

func Tasks(verbose bool, limit int, requests chan chan Task, new <-chan Task, state <-chan interface{}) {
	var taskList = make([]Task, 0)
	var mutex = &sync.Mutex{}

	//output loop
	go func() {
		for {
			req := <-requests
			mutex.Lock()
			if len(taskList) == 0 {
				req <- Task{}
			} else {
				req <- taskList[0]
				if verbose {
					fmt.Println("task " + taskList[0].String() + " is given to a worker")
				}
				taskList = taskList[1:]
			}
			mutex.Unlock()
		}
	}()

	//input loop
	go func() {
		for {
			task := <-new
			mutex.Lock()
			if len(taskList) >= limit {
				if verbose {
					fmt.Println("TASK STORAGE IS FULL")
				}
			} else {
				taskList = append(taskList, task)
			}
			mutex.Unlock()
		}
	}()

	//state loop
	for {
		<-state
		mutex.Lock()
		fmt.Println(taskList)
		mutex.Unlock()
	}
}

func Items(verbose bool, limit int, requests chan chan Item, new <-chan Item, state <-chan interface{}) {
	var itemList = make([]Item, 0)
	var mutex = &sync.Mutex{}

	go func() {
		for {
			req := <-requests
			mutex.Lock()
			if len(itemList) == 0 {
				req <- Item{}
			} else {
				req <- itemList[0]
				if verbose {
					fmt.Println("task " + itemList[0].String() + " is given to a worker")
				}
				itemList = itemList[1:]
			}
			mutex.Unlock()
		}
	}()

	go func() {
		for {
			item := <-new
			mutex.Lock()
			if len(itemList) >= limit {
				if verbose {
					fmt.Println("task storage is full")
				}
			} else {
				itemList = append(itemList, item)
			}
			mutex.Unlock()
		}
	}()

	for {
		<-state
		fmt.Println(itemList)
	}
}
