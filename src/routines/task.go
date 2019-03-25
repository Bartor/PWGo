package routines

import "errors"

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

func (t *Task) String() string {
	return "Task {" + string(t.Fst) + ", " + string(t.Snd) + "}"
}
