package routines

import (
	"fmt"
)

func taskInGuard(pass bool, channel chan Task) chan Task {
	if pass {
		return channel
	}
	return nil
}

func taskOutGuard(pass bool, channel chan chan Task) chan chan Task {
	if pass {
		return channel
	}
	return nil
}

func itemInGuard(pass bool, channel chan Item) chan Item {
	if pass {
		return channel
	}
	return nil
}

func itemOutGuard(pass bool, channel chan chan Item) chan chan Item {
	if pass {
		return channel
	}
	return nil
}

func Tasks(limit int, inTasks chan Task, outTasks chan chan Task, state <-chan interface{}) {
	var taskList = make([]Task, 0)

	for {
		select {
		case req := <-taskOutGuard(len(taskList) > 0, outTasks):
			var task = taskList[0]
			taskList = taskList[1:]
			req <- task
		case task := <-taskInGuard(len(taskList) < limit, inTasks):
			taskList = append(taskList, task)
		case <-state:
			fmt.Println(taskList)
		}
	}
}

func Items(limit int, inItems chan Item, outItems chan chan Item, state <-chan interface{}) {
	var itemList = make([]Item, 0)

	for {
		select {
		case req := <-itemOutGuard(len(itemList) > 0, outItems):
			var item = itemList[0]
			itemList = itemList[1:]
			req <- item
		case item := <-itemInGuard(len(itemList) < limit, inItems):
			itemList = append(itemList, item)
		case <-state:
			fmt.Println(itemList)
		}
	}

}
