package task

import (
	"github.com/kfchen81/beego/vanilla"
	_ "teamdo/business/account"
	b_task "teamdo/business/task"
	b_tunnel "teamdo/business/tunnel"
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
			"remark",
			"title",
		},
		"DELETE": []string{
			"id:int",
		},
	}
}

func (this *Task) Get()  {
	bCtx := this.GetBusinessContext()

	id, _ := this.GetInt("id")

	task := b_task.NewTaskRepository(bCtx).GetTaskById(id)
	data := b_task.NewEncodeTaskService(bCtx).Encode(task)

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

	tunnel := b_tunnel.NewTunnelRepository(bCtx).GetTunnelById(tunnelId)
	taskParams := b_tunnel.NewTaskParams(executorId, title, remark, priority, startDateStr, endDateStr)
	tunnel.AddTask(taskParams)

	response := vanilla.MakeResponse(vanilla.Map{})
	this.ReturnJSON(response)
}

func (this *Task) Post()  {
	bCtx := this.GetBusinessContext()

	id, _ := this.GetInt("id")
	remark := this.GetString("remark")
	title := this.GetString("title")

	task := b_task.NewTaskRepository(bCtx).GetTaskById(id)
	task.Update(remark, title)

	response := vanilla.MakeResponse(vanilla.Map{})
	this.ReturnJSON(response)
}

func (this *Task) Delete()  {
	bCtx := this.GetBusinessContext()

	id, _ := this.GetInt("id")
	task := b_task.NewTaskRepository(bCtx).GetTaskById(id)
	task.Delete()

	response := vanilla.MakeResponse(vanilla.Map{})
	this.ReturnJSON(response)
}