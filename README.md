# easy-crontab
做Go最简单的定时任务调度器

## 开箱即用
1. 在/job下按照示例新增任务
2. 在easy.go里面新增方法将任务调度器初始化
```go
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
```
3. 在easy.go里面TaskInit方法中加上你新增的定时任务
4. 执行go run easy.go 快速开始