package project

import (
	"github.com/kfchen81/beego/vanilla"
	b_project "teamdo/business/project"
)

type TaskStatus struct {
	vanilla.RestResource
}

func (this *TaskStatus) Resource() string {
	return "project.task_status"
}

func (this *TaskStatus) GetParameters() map[string][]string {
	return map[string][]string{
		"POST": []string{
			"id:int",
			"status:bool",
		},
	}
}

func (this *TaskStatus) Post()  {
	bCtx := this.GetBusinessContext()

	id, _ := this.GetInt("id")
	status, _ := this.GetBool("status")
	task := b_project.NewTaskRepository(bCtx).GetTaskById(id)
	task.UpdateStatus(status)

	response := vanilla.MakeResponse(vanilla.Map{})
	this.ReturnJSON(response)
}
