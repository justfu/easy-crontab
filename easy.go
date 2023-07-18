package easy

import (
	"easy-crontab/job"
	"github.com/reugn/go-quartz/quartz"
	"log"
	"sync"
)

func main() {
	(&CommonTask{}).Exec()
}

type CommonTask struct {
	CronTrigger *quartz.CronTrigger
	Job         quartz.Job
}

var TaskSlice []CommonTask

func (receiver *CommonTask) Exec() {
	var wg sync.WaitGroup
	wg.Add(1)
	sched := quartz.NewStdScheduler()
	sched.Start()
	log.Println("定时任务初始化开始!")
	receiver.TaskInit()
	log.Println("定时任务初始化成功!")
	for _, task := range TaskSlice {
		sched.ScheduleJob(task.Job, task.CronTrigger)
	}
	wg.Wait()
	sched.Stop()
}

// 定时任务初始化
func (receiver *CommonTask) TaskInit() {
	receiver.GetExample1()
	receiver.GetExample2()
}

func (receiver *CommonTask) GetExample1() {
	cronTrigger, err := quartz.NewCronTrigger("0 */1 * * * ?")
	if err != nil {
		panic(err)
	}
	cronJob := &job.Example1{Desc: "测试"}
	TaskSlice = append(TaskSlice, CommonTask{
		CronTrigger: cronTrigger,
		Job:         cronJob,
	})
}

func (receiver *CommonTask) GetExample2() {
	cronTrigger, err := quartz.NewCronTrigger("0 */1 * * * ?")
	if err != nil {
		panic(err)
	}
	cronJob := &job.Example2{Desc: "Example2"}
	TaskSlice = append(TaskSlice, CommonTask{
		CronTrigger: cronTrigger,
		Job:         cronJob,
	})
}
