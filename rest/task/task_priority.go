package task

import (
	"github.com/kfchen81/beego/vanilla"
	b_task "teamdo/business/task"
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
			"priority: sting",
		},
	}
}

func (this *TaskPriority) Post()  {
	bCtx := this.GetBusinessContext()

	id, _ := this.GetInt("task_id")
	priority := this.GetString("priority")
	task := b_task.NewTaskRepository(bCtx).GetTaskById(id)
	task.UpdatePriority(priority)

	response := vanilla.MakeResponse(vanilla.Map{})
	this.ReturnJSON(response)
}
