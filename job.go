package dispatch

import (
	"fmt"
	"time"
)

// Job is an interface
type Job interface {
	Do() error           // sync
	Description() string //
}

// DisplayJob is an imp of Job interface
type DisplayJob struct {
	Title string
}

// Do Do
func (job *DisplayJob) Do() error {
	time.Sleep(2000 * time.Millisecond)
	fmt.Println(job.Description())
	return nil
}

// Description Description
func (job *DisplayJob) Description() string {
	return "Display job " + job.Title
}

// OutputJob is an imp of Job interface
type OutputJob struct {
	Output string
}

// Do sleep for 5 seconds
func (job OutputJob) Do() error {
	time.Sleep(1000 * time.Millisecond)
	fmt.Println(job.Description())
	return nil
}

// Description Description
func (job OutputJob) Description() string {
	return "Output job " + job.Output
}
