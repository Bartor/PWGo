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
		Fst:    rand.Int(),
		Snd:    rand.Int(),
		Opr:    ops[rand.Int()%3],
		Broken: false,
	}
}

func Ceo(verbose bool, taskList chan Task, lo time.Duration, hi time.Duration) {
	for {
		var task = newTask()
		taskList <- task
		if verbose {
			fmt.Println("===[CEO] added a new task " + task.String())
		}

		var delay = time.Duration(rand.Int()%int(lo) + int(hi))
		time.Sleep(delay)
	}
}
