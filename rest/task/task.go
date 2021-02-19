package task

import (
	"github.com/kfchen81/beego/vanilla"
	b_project "teamdo/business/project"
	b_task "teamdo/business/task"
)

type Task struct {
	vanilla.RestResource
}

func (this *Task) Resource() string {
	return "task.task"
}

func (this *Task) GetParameters() map[string][]string {
	return map[string][]string{
		"GET": []string{"id:int", "?with_options:json"},
		"PUT": []string{
			"content:string",
			"project_id:int",
			"?status:int",
			"?user_id:int",
			"?parent_task_id:int",
			"?lane_id:int",
		},
		"POST": []string{
			"id:int",
			"project_id",
			"?content:string",
			"?status:int",
			"?lane_id:int",
			"?user_id:int",
		},
		"DELETE": []string{"id:int", "project_id"},
	}
}

func (this *Task) Get() {
	id, _ := this.GetInt("id")
	bCtx := this.GetBusinessContext()

	repository := b_task.NewTaskRepository(bCtx)
	task := repository.GetTaskById(id)
	if task == nil {
		panic(vanilla.NewBusinessError("task_not_exist", "任务不存在"))
	}
	b_task.NewFillTaskService(bCtx).FillOne(task, this.GetFillOptions("with_options"))

	encodeService := b_task.NewEncodeTaskService(bCtx)
	respData := encodeService.Encode(task)
	response := vanilla.MakeResponse(respData)
	this.ReturnJSON(response)
}

func (this *Task) Put() {
	content := this.GetString("content")
	pid, _ := this.GetInt("project_id")
	status, _ := this.GetInt("status")
	user_id, _ := this.GetInt("user_id")
	parent_task_id, _ := this.GetInt("parent_task_id")
	lane_id, _ := this.GetInt("lane_id")
	project_id, _ := this.GetInt("project_id")
	bCtx := this.GetBusinessContext()

	project := b_project.NewProjectRepository(bCtx).GetProjectById(pid)
	project.AuthorityVerify()

	task := b_task.NewTask(bCtx, content, status, user_id, parent_task_id, lane_id, project_id)
	response := vanilla.MakeResponse(vanilla.Map{
		"id": task.Id,
	})
	this.ReturnJSON(response)
}

func (this *Task) Delete() {
	//id, _ := this.GetInt("id")
	//pid, _ := this.GetInt("project_id")
	//bCtx := this.GetBusinessContext()
	//
	//project := b_project.NewProjectRepository(bCtx).GetProjectById(pid)
	//project.AuthorityVerify()
	//task := b_task.NewTaskRepository(bCtx).GetTaskById(id)
	//task.Delete()
	//
	//response := vanilla.MakeResponse(vanilla.Map{
	//	"id": id,
	//})
	//this.ReturnJSON(response)
}

func (this *Task) Post() {
	//id, _ := this.GetInt("id")
	//pid, _ := this.GetInt("project_id")
	//status, _ := this.GetInt("status")
	//content := this.GetString("content")
	//lane_id, _ := this.GetInt("lane_id")
	//user_id, _ := this.GetInt("user_id")
	//bCtx := this.GetBusinessContext()
	//
	//project := b_project.NewProjectRepository(bCtx).GetProjectById(pid)
	//project.AuthorityVerify()
	//
	//repository := b_lane.NewLaneRepository(bCtx)
	//lane := repository.GetLaneById(id)
	//lane.Update(name)
	//
	//response := vanilla.MakeResponse(vanilla.Map{})
	//this.ReturnJSON(response)
}
