package cron

import (
	"github.com/kfchen81/beego"
	"github.com/kfchen81/beego/vanilla/cron"
)

type demoTask struct {
	cron.Task
}


func (this *demoTask) Run(taskCtx *cron.TaskContext) error {
	beego.Info("[demo_task] run...")
	return nil
}

func NewDemoTask() *demoTask{
	task := new(demoTask)
	task.Task = cron.NewTask("demo_task")
	return task
}

func init() {
	//task := NewDemoTask()
	//cron.RegisterTask(task, "0/5 * * * * *")
}