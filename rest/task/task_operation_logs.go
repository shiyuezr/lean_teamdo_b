package task

import (
	"github.com/kfchen81/beego/vanilla"
	b_task_operation_log "teamdo/business/task_operation_log"
)

type TaskOperationLogs struct {
	vanilla.RestResource
}

func (this *TaskOperationLogs) Resource() string  {
	return "task.task_operation_logs"
}

func (this *TaskOperationLogs) GetParameters() map[string][]string {
	return map[string][]string{
		"GET": {"?task_id:int", "?project_id:int", "?filters:json", "?with_options:json"},
	}
}

func (this *TaskOperationLogs) Get() {
	projectId, _ := this.GetInt("project_id")
	taskId, _ := this.GetInt("task_id")
	bCtx := this.GetBusinessContext()
	filters := vanilla.ConvertToBeegoOrmFilter(this.GetFilters())
	logs := b_task_operation_log.NewTaskOperationLogRepository(bCtx).GetLogsByProjectAndTask(projectId, taskId, filters)

	b_task_operation_log.NewFillTaskOperationLogService(bCtx).Fill(logs, this.GetFillOptions("with_options"))
	response := vanilla.MakeResponse(b_task_operation_log.NewEncodeTaskOperationLogService(bCtx).EncodeMany(logs))
	this.ReturnJSON(response)
}