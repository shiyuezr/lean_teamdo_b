package task

import (
	"github.com/kfchen81/beego/vanilla"
	_ "teamdo/business/account"
	task2 "teamdo/business/task"
)

type TaskTitle struct {
	vanilla.RestResource
}

func (this *TaskTitle) Resource() string {
	return "project.task_title"
}

func (this *TaskTitle) GetParameters() map[string][]string {
	return map[string][]string{
		"POST": []string{
			"task_id:int",
			"title",
		},
	}
}

func (this *TaskTitle) Post()  {
	bCtx := this.GetBusinessContext()

	id, _ := this.GetInt("task_id")
	title := this.GetString("title")
	task := task2.NewTaskRepository(bCtx).GetTaskById(id)

	task.UpdateTitle(title)

	response := vanilla.MakeResponse(vanilla.Map{})
	this.ReturnJSON(response)
}
