package task

import (
	"github.com/kfchen81/beego/vanilla"
	b_task "teamdo/business/task"
)

type UnFinishTask struct {
	vanilla.RestResource
}

func (this *UnFinishTask) Resource() string {
	return "project.unFinish_task"
}

func (this *UnFinishTask) GetParameters() map[string][]string {
	return map[string][]string{
		"PUT": []string{
			"task_id:int",
		},
	}
}

func (this *UnFinishTask) Put()  {
	bCtx := this.GetBusinessContext()

	id, _ := this.GetInt("task_id")
	task := b_task.NewTaskRepository(bCtx).GetTaskById(id)
	task.UnFinishTask()

	response := vanilla.MakeResponse(vanilla.Map{})
	this.ReturnJSON(response)
}
