package dispatch

// Dispatcher Dispatcher
type Dispatcher struct {
	JobQueue chan Job     // accept incoming job
	WorkPool chan *Worker // the pool of workers
}

// NewDispatcher create a new *Dispatcher
func NewDispatcher(maxWorkerCount int) *Dispatcher {
	pool := make(chan *Worker, maxWorkerCount)
	return &Dispatcher{
		WorkPool: pool,
		JobQueue: make(chan Job)}
}

// Run the loop
func (dispatcher *Dispatcher) Run() {
	count := cap(dispatcher.WorkPool)
	for i := 0; i < count; i++ {
		worker := NewWorker(dispatcher)
		worker.Start()
	}

	go dispatcher.dispatch()
}

// dispatch job
func (dispatcher *Dispatcher) dispatch() {
	for {
		select {
		case job := <-dispatcher.JobQueue:
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
