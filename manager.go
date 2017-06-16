package dispatch

import "errors"

// ErrorManagerNotStart ErrorNotStartYet
var ErrorManagerNotStart = errors.New("Manager is not not running")

// Manager Core
type Manager struct {
	workPool    chan *Worker // the pool of workers
	jobQueue    chan Job     // accept incoming job
	stopChannel chan bool    // 停止
	isReady     bool         // 是处于可用状态
}

// NewManager create a new *Manager
func NewManager(maxWorkerCount int) *Manager {
	//
	manager := &Manager{
		workPool:    make(chan *Worker, maxWorkerCount),
		jobQueue:    make(chan Job),
		stopChannel: make(chan bool),
		isReady:     false}
	//
	manager.Setup()
	return manager
}

// Setup 创建workers
func (manager *Manager) Setup() {
	count := cap(manager.workPool)
	for i := 0; i < count; i++ {
		worker := NewWorker(manager)
		worker.Start()
	}
}

// Listen 开始接受任务
func (manager *Manager) Listen() {
	go manager.dispatch()
	manager.isReady = true
}

// Accept 接收到新的任务
func (manager *Manager) Accept(job Job) error {
	if manager.IsReady() {
		manager.jobQueue <- job
		return nil
	}
	return ErrorManagerNotStart
}

// Stop 停止接受任务
func (manager *Manager) Stop() {
	manager.isReady = false
	manager.stopChannel <- true
}

// IsReady 是否可用
func (manager *Manager) IsReady() bool {
	return manager.isReady
}

// dispatch 派发任务
func (manager *Manager) dispatch() {
	for {
		select {
		case job := <-manager.jobQueue:
			go func(job Job) {
				// try to obtain a worker job channel that is available.
				// this will block until a worker is idle
				worker := <-manager.workPool
				// dispatch the job to the worker job channel
				worker.Receive(job)
			}(job)
		case <-manager.stopChannel:
			return
		}
	}
}
