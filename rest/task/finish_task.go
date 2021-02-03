package task

import (
	"github.com/kfchen81/beego/vanilla"
	_ "teamdo/business/account"
	task2 "teamdo/business/task"
)

type FinishTask struct {
	vanilla.RestResource
}

func (this *FinishTask) Resource() string {
	return "project.task_status"
}

func (this *FinishTask) GetParameters() map[string][]string {
	return map[string][]string{
		"PUT": []string{
			"task_id:int",
		},
		"DELETE": []string{
			"task_id:int",
		},
	}
}

func (this *FinishTask) Put()  {
	bCtx := this.GetBusinessContext()

	id, _ := this.GetInt("task_id")
	task := task2.NewTaskRepository(bCtx).GetTaskById(id)
	task.CompleteTask()

	response := vanilla.MakeResponse(vanilla.Map{})
	this.ReturnJSON(response)
}

func (this *FinishTask) Delete()  {
	bCtx := this.GetBusinessContext()

	id, _ := this.GetInt("id")
	task := task2.NewTaskRepository(bCtx).GetTaskById(id)
	task.FailToFinishTask()

	response := vanilla.MakeResponse(vanilla.Map{})
	this.ReturnJSON(response)
}
