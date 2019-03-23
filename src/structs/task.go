package structs

import "errors"

type Task struct {
	Fst int
	Snd int
	Opr int
}

func (t *Task) ResolveTask() (int, error) {
	switch t.Opr {
	case 0:
		return t.Fst + t.Snd, nil
	case 1:
		return t.Fst - t.Snd, nil
	case 2:
		return t.Fst * t.Snd, nil
	default:
		return 0, errors.New("wrong op")
	}
}
