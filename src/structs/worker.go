package structs

type Worker struct {
	Delay int
	Tasks chan Task
	Results chan int
}

