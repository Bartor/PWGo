package routines

import (
	"errors"
	"strconv"
	"time"
)

type WorkerConfig struct {
	Id      int
	Verbose bool
	Delay   time.Duration
	Patient bool
}

type Item struct {
	Value int
}

type Task struct {
	Fst int
	Snd int
	Opr func(int, int) int
}

func (t *Task) ResolveTask() (int, error) {
	if t.Opr == nil {
		return 0, errors.New("no operation")
	} else {
		return t.Opr(t.Fst, t.Snd), nil
	}
}

func (t Task) String() string {
	return "Task {" + strconv.Itoa(t.Fst) + ", " + strconv.Itoa(t.Snd) + "}"
}

func (i Item) String() string {
	return "Item {" + strconv.Itoa(i.Value) + "}"
}
