package dispatch

// Jobable Jobable
type Jobable interface {
	Do()
}

// Job Job
type Job struct {
	Title       string
	Description string
}

// JobQueue 全局JobQueue
var JobQueue = make(chan Job)
