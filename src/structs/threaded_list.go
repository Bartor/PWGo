package structs

type ThreadList interface {
	Start()
}

type ThreadListTask struct {
	List []Task
	In   chan<- Task
	Out  <-chan Task
}

type ThreadListInt struct {
	List []int
	In   chan<- int
	Out  <-chan int
}

func (list *ThreadListTask) Start() chan interface{} {
	//bad idea
	for {
		select {
		case newMsg := <-list.Out:
			list.List = append(list.List, newMsg)
		case list.In <- list.List[0]:
			println("struct!")
			list.List = list.List[1:]
		default:

		}
	}
}

func (list *ThreadListInt) Start() chan interface{} {
	//bad idea
	for {
		select {
		case newMsg := <-list.Out:
			list.List = append(list.List, newMsg)
		case list.In <- list.List[0]:
			println(list.List[0])
			list.List = list.List[1:]
		default:

		}
	}
}
