package structs

type ThreadList struct {
	List []interface{}
	In   chan interface{}
	Out  chan interface{}
}

func (list *ThreadList) Start() chan interface{} {
	for {
		select {
		case newMsg := <-list.In:
			list.List = append(list.List, newMsg)
		case list.Out <- list.List[0]:
			list.List = list.List[1:]
		}
	}
}
