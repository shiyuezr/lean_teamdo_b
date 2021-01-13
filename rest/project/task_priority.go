package project

import (
	"github.com/kfchen81/beego/vanilla"
	b_project "teamdo/business/project"
)

type TaskPriority struct {
	vanilla.RestResource
}

func (this *TaskPriority) Resource() string {
	return "project.task_priority"
}

func (this *TaskPriority) GetParameters() map[string][]string {
	return map[string][]string{
		"POST": []string{
			"id:int",
			"priority",
		},
	}
}

func (this *TaskPriority) Post()  {
	bCtx := this.GetBusinessContext()

	id, _ := this.GetInt("id")
	priority := this.GetString("priority")
	task := b_project.NewTaskRepository(bCtx).GetTaskById(id)
	task.UpdatePriority(priority)

	response := vanilla.MakeResponse(vanilla.Map{})
	this.ReturnJSON(response)
}
