package project

import (
	"github.com/kfchen81/beego/vanilla"
	b_project "teamdo/business/project"
)

type Task struct {
	vanilla.RestResource
}

func (this *Task) Resource() string {
	return "project.task"
}

func (this *Task) GetParameters() map[string][]string {
	return map[string][]string{
		"GET": []string{"id:int"},
		"PUT": []string{
			"manager_id:int",
			"tunnel_id:int",
			"?executor_id:int",
			"start_date",
			"end_date",
			"title",
			"remark",
			"priority",
		},
		"POST": []string{
			"id:int",
			"title",
			"remark",
		},
		"DELETE": []string{
			"id:int",
		},
	}
}

func (this *Task) Get()  {
	bCtx := this.GetBusinessContext()

	id, _ := this.GetInt("id")

	task := b_project.NewTaskRepository(bCtx).GetTaskById(id)
	data := b_project.NewEncodeTaskService(bCtx).Encode(task)

	response := vanilla.MakeResponse(vanilla.Map{
		"task": data,
	})
	this.ReturnJSON(response)
}

func (this *Task) Put()  {
	bCtx := this.GetBusinessContext()

	tunnelId, _ := this.GetInt("tunnel_id")
	executorId, _ := this.GetInt("executor_id", 0)
	title := this.GetString("title")
	remark := this.GetString("remark")
	priority := this.GetString("priority")
	startDateStr := this.GetString("start_date")
	endDateStr := this.GetString("end_date")

	tunnel := b_project.NewTunnelRepository(bCtx).GetTunnelById(tunnelId)
	tunnel.AddTask(executorId, title, remark, priority, startDateStr, endDateStr)

	response := vanilla.MakeResponse(vanilla.Map{})
	this.ReturnJSON(response)
}

func (this *Task) Post()  {
	bCtx := this.GetBusinessContext()

	id, _ := this.GetInt("id")
	title := this.GetString("title", "")
	remark := this.GetString("remark", "")

	task := b_project.NewTaskRepository(bCtx).GetTaskById(id)
	task.Update(title, remark)

	response := vanilla.MakeResponse(vanilla.Map{})
	this.ReturnJSON(response)
}

func (this *Task) Delete()  {
	bCtx := this.GetBusinessContext()

	id, _ := this.GetInt("id")
	task := b_project.NewTaskRepository(bCtx).GetTaskById(id)
	task.Delete()

	response := vanilla.MakeResponse(vanilla.Map{})
	this.ReturnJSON(response)
}