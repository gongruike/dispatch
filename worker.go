package dispatch

// Worker 负责完成job
type Worker struct {
	manager     *Manager      //
	jobChannel  chan Job      // 任务队列，从任务队列里获取任务，每一个worker都有一个单独的jobChannel
	stopChannel chan struct{} // 停止工作
}

// NewWorker 创建一个新的worker
func NewWorker(manager *Manager) *Worker {
	return &Worker{
		manager:     manager,
		jobChannel:  make(chan Job),
		stopChannel: make(chan struct{})}
}

// Start 开始接受job
func (worker *Worker) Start() {
	go func() {
		for {
			// tell the manager that I'm ready to work now
			// it will not block because it's a buffer channel
			worker.manager.register(worker)
			select {
			// blocked & waiting for new job
			case job := <-worker.jobChannel:
				job.Do()
			case <-worker.stopChannel:
				return
			}
		}
	}()
}

// Receive 收到新的任务
func (worker *Worker) Receive(job Job) {
	worker.jobChannel <- job
}

// Stop 停止工作
func (worker *Worker) Stop() {
	// send a message
	worker.stopChannel <- struct{}{}
}
