package task

import (
	"github.com/kfchen81/beego/vanilla"
	b_account "teamdo/business/account"
	b_project "teamdo/business/project"
	b_task "teamdo/business/task"
	b_task_operation_log "teamdo/business/task_operation_log"
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
			"lane_id:int",
			"?operator_id:int",
			"?parent_task_id:int",
		},
		"POST": []string{
			"id:int",
			"project_id",
			"?content:string",
			"?status:int",
			"?lane_id:int",
			"?operator_id:int",
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
	projectId, _ := this.GetInt("project_id")
	laneId, _ := this.GetInt("lane_id")
	operatorId, _ := this.GetInt("operator_id")
	parentTaskId, _ := this.GetInt("parent_task_id")
	bCtx := this.GetBusinessContext()
	userId := b_account.GetUserFromContext(bCtx).Id

	project := b_project.NewProjectRepository(bCtx).GetProjectById(projectId)
	project.AuthorityVerify()

	task := b_task.NewTask(bCtx, content, operatorId, parentTaskId, laneId, projectId)
	b_task_operation_log.NewTaskOperationLog(bCtx, "创建", task.Id, userId)
	response := vanilla.MakeResponse(vanilla.Map{
		"id": task.Id,
	})
	this.ReturnJSON(response)
}

func (this *Task) Delete() {
	id, _ := this.GetInt("id")
	pid, _ := this.GetInt("project_id")
	bCtx := this.GetBusinessContext()
	userId := b_account.GetUserFromContext(bCtx).Id

	project := b_project.NewProjectRepository(bCtx).GetProjectById(pid)
	// 只有任务的管理员才可以删除项目下的任务
	project.AuthorityVerify()
	task := b_task.NewTaskRepository(bCtx).GetTaskById(id)
	task.Delete()
	b_task_operation_log.NewTaskOperationLog(bCtx, "删除", task.Id, userId)
	response := vanilla.MakeResponse(vanilla.Map{
		"id": id,
	})
	this.ReturnJSON(response)
}

func (this *Task) Post() {
	id, _ := this.GetInt("id")
	pid, _ := this.GetInt("project_id")
	status, _ := this.GetInt("status")
	content := this.GetString("content")
	laneId, _ := this.GetInt("lane_id")
	operatorId, _ := this.GetInt("operator_id")
	bCtx := this.GetBusinessContext()
	userId := b_account.GetUserFromContext(bCtx).Id

	project := b_project.NewProjectRepository(bCtx).GetProjectById(pid)
	project.AuthorityVerify()

	repository := b_task.NewTaskRepository(bCtx)
	task := repository.GetTaskById(id)
	task.Update(status, content, laneId, operatorId)
	b_task_operation_log.NewTaskOperationLog(bCtx, "修改", task.Id, userId)

	response := vanilla.MakeResponse(vanilla.Map{})
	this.ReturnJSON(response)
}
