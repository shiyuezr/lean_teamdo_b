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
			"date:json",
			"title:string",
			"remark:string",
			"priority:string",
		},
		"POST": []string{
			"id:int",
			"title:string",
			"remark:string",
			"priority:string",
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

// 所有的创建,修改,删除,必须是管理员, 这个权限未加
func (this *Task) Put()  {
	bCtx := this.GetBusinessContext()
	//managerId, _ := this.GetInt("manager_id")
	projectName := this.GetString("project_name")
	tunnelId, _ := this.GetInt("tunnel_id")
	executorId, _ := this.GetInt("executor_id")
	title := this.GetString("title")
	status, _ := this.GetBool("status")
	remark := this.GetString("remark")
	priority := this.GetString("priority")

	b_project.NewTask(bCtx, title, executorId, tunnelId, status, remark, priority, projectName)
	response := vanilla.MakeResponse(vanilla.Map{})
	this.ReturnJSON(response)
}

func (this *Task) Post()  {
	bCtx := this.GetBusinessContext()

	id, _ := this.GetInt("id")
	title := this.GetString("title")
	status, _ := this.GetBool("status")
	remark := this.GetString("remark")
	priority := this.GetString("priority")

	task := b_project.NewTaskRepository(bCtx).GetTaskById(id)
	task.Update(title, status, remark, priority)

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