package routines

import (
	"strconv"
	"time"
)

type WorkerConfig struct {
	Id      int
	Verbose bool
	Delay   time.Duration
	Timeout time.Duration
	Patient bool
}

type MachineConfig struct {
	Id      int
	Verbose bool
	Delay   time.Duration
}

type Item struct {
	Value int
}

type Task struct {
	Fst int
	Snd int
	Opr func(int, int) int
	Res int
}

func (t *Task) ResolveTask() {
	if t.Opr == nil {
		t.Res = 0
	} else {
		t.Res = t.Opr(t.Fst, t.Snd)
	}
}

func (t Task) String() string {
	return "Task {" + strconv.Itoa(t.Fst) + ", " + strconv.Itoa(t.Snd) + "}"
}

func (i Item) String() string {
	return "Item {" + strconv.Itoa(i.Value) + "}"
}
