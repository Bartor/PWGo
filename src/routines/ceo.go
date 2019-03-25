package routines

import (
	"fmt"
	"math/rand"
	"time"
)

func add(a int, b int) int { return a + b }
func mul(a int, b int) int { return a * b }
func sub(a int, b int) int { return a - b }

func newTask() Task {
	var ops = [3]func(int, int) int{add, mul, sub}
	return Task{
		Fst: rand.Int(),
		Snd: rand.Int(),
		Opr: ops[rand.Int()%3],
	}
}

func Ceo(verbose bool, taskList chan<- Task, delay time.Duration) {
	for {
		var task = newTask()
		taskList <- task
		if verbose {
			fmt.Println("ceo added a new task " + task.String())
		}

		time.Sleep(delay)
	}
}