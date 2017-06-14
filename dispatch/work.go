package dispatch

import "time"
import "fmt"

// Worker Worker
type Worker struct {
	//
	JobChannel  chan Job     // 任务队列，从任务队列里获取任务，每一个worker都有一个单独的JobQueueChannel
	WorkerPool  chan *Worker // 用来笑死通讯，当前worker已经处于可用状态
	QuitChannel chan bool    // 退出
}

// NewWorker NewWorker
func NewWorker(workerPool chan *Worker) *Worker {
	return &Worker{
		WorkerPool:  workerPool,
		JobChannel:  make(chan Job),
		QuitChannel: make(chan bool)}
}

// Start worker loop
func (worker *Worker) Start() {
	go func() {
		for {
			// tell the manager, the worker is ready now
			worker.WorkerPool <- worker

			select {
			case job := <-worker.JobChannel:
				time.Sleep(5000 * time.Millisecond)
				fmt.Println(job.Title + job.Description)
			case <-worker.QuitChannel:
				return
			}
		}
	}()
}

// Do the job
func (worker *Worker) Do() {

}
