package task

import (
	"github.com/kfchen81/beego/vanilla"
	b_task "teamdo/business/task"
)

type FinishTask struct {
	vanilla.RestResource
}

func (this *FinishTask) Resource() string {
	return "project.finish_task"
}

func (this *FinishTask) GetParameters() map[string][]string {
	return map[string][]string{
		"PUT": []string{
			"task_id:int",
		},
	}
}

func (this *FinishTask) Put()  {
	bCtx := this.GetBusinessContext()

	id, _ := this.GetInt("task_id")
	task := b_task.NewTaskRepository(bCtx).GetTaskById(id)
	task.FinishTask()

	response := vanilla.MakeResponse(vanilla.Map{})
	this.ReturnJSON(response)
}
