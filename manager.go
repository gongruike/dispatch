package dispatch

// Manager Manager
type Manager struct {
	jobQueue    chan Job     // accept incoming job
	workPool    chan *Worker // the pool of workers
	stopChannel chan bool    // stop

}

// NewManager create a new *Manager
func NewManager(maxWorkerCount int) *Manager {
	pool := make(chan *Worker, maxWorkerCount)
	return &Manager{
		workPool:    pool,
		jobQueue:    make(chan Job),
		stopChannel: make(chan bool)}
}

// Listen 初始化Worker并开始接受任务
func (manager *Manager) Listen() {
	count := cap(manager.workPool)
	for i := 0; i < count; i++ {
		// create workers
		worker := NewWorker(manager)
		worker.Start()
	}
	//
	go manager.dispatch()
}

// Accept 接收到新的任务
func (manager *Manager) Accept(job Job) {
	manager.jobQueue <- job
}

// Stop 停止
func (manager *Manager) Stop() {
	manager.stopChannel <- true
}

// dispatch job
func (manager *Manager) dispatch() {
	for {
		select {
		case job := <-manager.jobQueue:
			go func(job Job) {
				// try to obtain a worker job channel that is available.
				// this will block until a worker is idle
				worker := <-manager.workPool
				// dispatch the job to the worker job channel
				worker.jobChannel <- job
			}(job)
		case <-manager.stopChannel:
			return
		}
	}
}
