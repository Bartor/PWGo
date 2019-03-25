package routines

func Tasks(verbose bool, limit int, requests <-chan GetRequestTask, new <-chan Task, state <-chan interface{}) {
	var taskList = make([]Task, 0)
	for {
		select {
		case req := <-requests:
			if len(taskList) != 0 {
				req.Response <- Task{}
			}
		}
	}
}
