package structs

import (
	"math/rand"
	"time"
)

type Ceo struct {
	LoLimit int
	HiLimit int
	Tasks   chan<- Task
}

func (ceo *Ceo) newTask() {
	t := Task{rand.Int(), rand.Int(), rand.Int() % 3}
	ceo.Tasks <- t
}

func (ceo *Ceo) Start() {
	println("starting ceo")
	for {
		time.Sleep(time.Duration(rand.Int()%ceo.HiLimit+ceo.LoLimit) * time.Second)
		ceo.newTask()
	}
}
