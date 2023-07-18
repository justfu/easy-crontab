package job

import (
	"easy-crontab/errorcollect"
	"fmt"
	"github.com/reugn/go-quartz/quartz"
)

// PrintJob implements the quartz.Job interface.
type Example1 struct {
	Desc string
}

// Description returns the description of the PrintJob.
func (pj *Example1) Description() string {
	return pj.Desc
}

// Key returns the unique PrintJob key.
func (pj *Example1) Key() int {
	return quartz.HashCode(pj.Description())
}

// Execute is called by a Scheduler when the Trigger associated with this job fires.
func (pj *Example1) Execute() {
	defer errorcollect.ErrorCollect(pj.Desc)
	fmt.Println("example1")
}
