package task

import (
	"github.com/kfchen81/beego/vanilla"
	task2 "teamdo/business/task"
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
			"task_id:int",
			"priority: int",
		},
	}
}

func (this *TaskPriority) Post()  {
	bCtx := this.GetBusinessContext()

	id, _ := this.GetInt("task_id")
	priority := this.GetString("priority")
	task := task2.NewTaskRepository(bCtx).GetTaskById(id)
	task.UpdatePriority(priority)

	response := vanilla.MakeResponse(vanilla.Map{})
	this.ReturnJSON(response)
}
