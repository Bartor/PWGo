package routines

import (
	"fmt"
	"sync"
)

func Tasks(verbose bool, limit int, requests chan chan Task, new <-chan Task, state <-chan interface{}) {
	var taskList = make([]Task, 0)
	var reqQueue = make([]chan Task, 0)
	var taskMutex = &sync.Mutex{}
	var reqsMutex = &sync.Mutex{}

	//output loop
	go func() {
		for {
			req := <-requests
			taskMutex.Lock()
			if len(taskList) == 0 {
				reqsMutex.Lock()
				reqQueue = append(reqQueue, req)
				reqsMutex.Unlock()
			} else {
				req <- taskList[0]
				if verbose {
					fmt.Println("task " + taskList[0].String() + " is given to a worker")
				}
				taskList = taskList[1:]
			}
			taskMutex.Unlock()
		}
	}()

	//input loop
	go func() {
		for {
			task := <-new
			taskMutex.Lock()
			if len(taskList) >= limit {
				reqsMutex.Lock()
				if len(reqQueue) > 0 {
					reqQueue[0] <- taskList[0]
					reqQueue = reqQueue[1:]
					taskList = taskList[1:]
					taskList = append(taskList, task)
				} else if verbose {
					fmt.Println("TASK STORAGE IS FULL")
				}
				reqsMutex.Unlock()
			} else {
				reqsMutex.Lock()
				taskList = append(taskList, task)
				if len(reqQueue) > 0 {
					reqQueue[0] <- taskList[0]
					reqQueue = reqQueue[1:]
					taskList = taskList[1:]
				}
				reqsMutex.Unlock()
			}
			taskMutex.Unlock()
		}
	}()

	//state loop
	for {
		<-state
		taskMutex.Lock()
		fmt.Println(taskList)
		taskMutex.Unlock()
	}
}

func Items(verbose bool, limit int, requests chan chan Item, new <-chan Item, state <-chan interface{}) {
	var itemList = make([]Item, 0)
	var reqQueue = make([]chan Item, 0)
	var reqsMutex = &sync.Mutex{}
	var taskMutex = &sync.Mutex{}

	go func() {
		for {
			req := <-requests
			taskMutex.Lock()
			if len(itemList) == 0 {
				reqsMutex.Lock()
				reqQueue = append(reqQueue, req)
				reqsMutex.Unlock()
			} else {
				req <- itemList[0]
				if verbose {
					fmt.Println("item " + itemList[0].String() + " is given to a client")
				}
				itemList = itemList[1:]
			}
			taskMutex.Unlock()
		}
	}()

	go func() {
		for {
			item := <-new
			taskMutex.Lock()
			if len(itemList) >= limit {
				if verbose {
					fmt.Println("task storage is full")
				}
			} else {
				itemList = append(itemList, item)
			}
			taskMutex.Unlock()
		}
	}()

	for {
		item := <-new
		taskMutex.Lock()
		if len(itemList) >= limit {
			reqsMutex.Lock()
			if len(reqQueue) > 0 {
				reqQueue[0] <- itemList[0]
				reqQueue = reqQueue[1:]
				itemList = itemList[1:]
				itemList = append(itemList, item)
			} else if verbose {
				fmt.Println("ITEM STORAGE IS FULL")
			}
			reqsMutex.Unlock()
		} else {
			reqsMutex.Lock()
			itemList = append(itemList, item)
			if len(reqQueue) > 0 {
				reqQueue[0] <- itemList[0]
				reqQueue = reqQueue[1:]
				itemList = itemList[1:]
			}
			reqsMutex.Unlock()
		}
		taskMutex.Unlock()
	}
}
