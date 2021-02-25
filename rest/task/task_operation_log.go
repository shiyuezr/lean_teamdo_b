package task

import (
	"fmt"
	"github.com/kfchen81/beego/vanilla"
	b_account "teamdo/business/account"
	b_task_operation_log "teamdo/business/task_operation_log"
)

type TaskOperationLog struct {
	vanilla.RestResource
}

func (this *TaskOperationLog) Resource() string {
	return "task.task_operation_log"
}

func (this *TaskOperationLog) GetParameters() map[string][]string {
	return map[string][]string{
		"GET": []string{"id:int", "?with_options:json"},
		"PUT": []string{
			"task_id:int",
			"content:string",
		},
	}
}

func (this *TaskOperationLog) Put() {
	content := this.GetString("content")
	task_id, _ := this.GetInt("task_id")
	bCtx := this.GetBusinessContext()
	uid := b_account.GetUserFromContext(bCtx).Id
	fmt.Println(uid)
	log := b_task_operation_log.NewTaskOperationLog(bCtx, content, task_id, uid)
	response := vanilla.MakeResponse(vanilla.Map{
		"id": log.Id,
	})
	this.ReturnJSON(response)
}

func (this *TaskOperationLog) Get() {
	id, _ := this.GetInt("id")
	bCtx := this.GetBusinessContext()

	repository := b_task_operation_log.NewTaskOperationLogRepository(bCtx)
	log := repository.GetTaskById(id)
	if log == nil {
		panic(vanilla.NewBusinessError("log_not_exist", "日志不存在"))
	}
	b_task_operation_log.NewFillTaskOperationLogService(bCtx).FillOne(log, this.GetFillOptions("with_options"))

	encodeService := b_task_operation_log.NewEncodeTaskOperationLogService(bCtx)
	respData := encodeService.Encode(log)
	response := vanilla.MakeResponse(respData)
	this.ReturnJSON(response)
}
