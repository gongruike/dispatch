package dispatch

// Dispatcher Dispatcher
type Dispatcher struct {
	WorkPool chan *Worker
}

// NewDispatcher create a new *ManaDispatcherger
func NewDispatcher(maxWorkerCount int) *Dispatcher {
	pool := make(chan *Worker, maxWorkerCount)
	return &Dispatcher{WorkPool: pool}
}

// Run the loop
func (dispatcher *Dispatcher) Run() {
	count := cap(dispatcher.WorkPool)
	for i := 0; i < count; i++ {
		worker := NewWorker(dispatcher.WorkPool)
		worker.Start()
	}

	go dispatcher.dispatch()
}

// dispatch job
func (dispatcher *Dispatcher) dispatch() {
	for {
		select {
		case job := <-JobQueue:
			go func(job Job) {
				// try to obtain a worker job channel that is available.
				// this will block until a worker is idle
				worker := <-dispatcher.WorkPool
				// dispatch the job to the worker job channel
				worker.JobChannel <- job
			}(job)
		}
	}
}
