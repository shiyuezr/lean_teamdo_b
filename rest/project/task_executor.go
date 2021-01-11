package project

import (
	"github.com/kfchen81/beego/vanilla"
	b_project "teamdo/business/project"
)

type TaskExecutor struct {
	vanilla.RestResource
}

func (this *TaskExecutor) Resource() string {
	return "project.task_executor"
}

func (this *TaskExecutor) GetParameters() map[string][]string {
	return map[string][]string{
		"POST": []string{
			"id:int",
			"executor_id:int",
		},
	}
}

func (this *TaskExecutor) Post()  {
	bCtx := this.GetBusinessContext()

	id, _ := this.GetInt("id")
	executorId, _ := this.GetInt("executor_id")
	task := b_project.NewTaskRepository(bCtx).GetTaskById(id)
	task.UpdateExecutor(executorId)

	response := vanilla.MakeResponse(vanilla.Map{})
	this.ReturnJSON(response)
}