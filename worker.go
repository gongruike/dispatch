package dispatch

/*
// Workerable Workerable
type Workerable interface {
	Start() error
	Stop()
}
	本来想把Worker抽象出来，现在看来不太可能了
*/

// Worker 负责完成job
type Worker struct {
	manager     *Manager  //
	jobChannel  chan Job  // 任务队列，从任务队列里获取任务，每一个worker都有一个单独的jobChannel
	stopChannel chan bool // 停止工作
}

// NewWorker 创建一个新的worker
func NewWorker(manager *Manager) *Worker {
	return &Worker{
		manager:     manager,
		jobChannel:  make(chan Job),
		stopChannel: make(chan bool)}
}

// Start 开始接受job
func (worker *Worker) Start() {
	go func() {
		for {
			// tell the dispatcher, the worker is ready now
			worker.manager.workPool <- worker
			select {
			// blocked & waiting for incoming job
			case job := <-worker.jobChannel:
				job.Do()
			case <-worker.stopChannel:
				return
			}
		}
	}()
}

// Stop 停止工作
func (worker *Worker) Stop() {
	// send a message
	worker.stopChannel <- true
}
