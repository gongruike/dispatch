package dispatch

// Manager Manager
type Manager struct {
	JobQueue chan Job     // accept incoming job
	workPool chan *Worker // the pool of workers
}

// NewManager create a new *Manager
func NewManager(maxWorkerCount int) *Manager {
	pool := make(chan *Worker, maxWorkerCount)
	return &Manager{
		workPool: pool,
		JobQueue: make(chan Job)}
}

// Dispatch the loop
func (manager *Manager) Dispatch() {
	count := cap(manager.workPool)
	for i := 0; i < count; i++ {
		worker := NewWorker(manager)
		worker.Start()
	}
	//
	go manager.dispatch()
}

// dispatch job
func (manager *Manager) dispatch() {
	for {
		select {
		case job := <-manager.JobQueue:
			go func(job Job) {
				// try to obtain a worker job channel that is available.
				// this will block until a worker is idle
				worker := <-manager.workPool
				// dispatch the job to the worker job channel
				worker.jobChannel <- job
			}(job)
		}
	}
}
