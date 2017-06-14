package dispatch

// Worker Worker
type Worker struct {
	JobChannel  chan Job    // 任务队列，从任务队列里获取任务，每一个worker都有一个单独的JobQueueChannel
	ADispatcher *Dispatcher // 用来消息通讯，当前worker已经处于可用状态
}

// NewWorker NewWorker
func NewWorker(dispatcher *Dispatcher) *Worker {
	return &Worker{
		ADispatcher: dispatcher,
		JobChannel:  make(chan Job)}
}

// Start worker loop
func (worker *Worker) Start() {
	go func() {
		for {
			// tell the dispatcher, the worker is ready now
			worker.ADispatcher.WorkPool <- worker
			select {
			// blocked & waiting for incoming job
			case job := <-worker.JobChannel:
				job.Do()
			}
		}
	}()
}
