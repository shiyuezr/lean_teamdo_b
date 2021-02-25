package task_operation_log

import (
	"context"
	"github.com/kfchen81/beego"
	"github.com/kfchen81/beego/vanilla"
	b_task "teamdo/business/task"
	m_task "teamdo/models/task"
)

type TaskOperationLogRepository struct {
	vanilla.RepositoryBase
}

func (this *TaskOperationLogRepository) GetByFilters(filters vanilla.Map) []*TaskOperationLog {
	qs := vanilla.GetOrmFromContext(this.Ctx).QueryTable(&m_task.TaskOperationLog{})
	if len(filters) > 0 {
		qs = qs.Filter(filters)
	}

	var dbModels []*m_task.TaskOperationLog
	_, err := qs.OrderBy("-id").All(&dbModels)
	if err != nil {
		beego.Error(err)
		return []*TaskOperationLog{}
	}
	taskOperationLogs := make([]*TaskOperationLog, 0, len(dbModels))
	for _, dbModel := range dbModels {
		taskOperationLogs = append(taskOperationLogs, NewTaskOperationLogFromDbModel(this.Ctx, dbModel))
	}
	return taskOperationLogs
}

func (this *TaskOperationLogRepository) GetTaskById(id int) *TaskOperationLog {
	filters := vanilla.Map{
		"id": id,
	}
	logs := this.GetByFilters(filters)
	if len(logs) == 0 {
		return nil
	}
	return logs[0]
}

func (this *TaskOperationLogRepository) GetLogsByProjectAndTask(projectId, taskId int, filters vanilla.Map) []*TaskOperationLog{
	if projectId != 0 {
		filtersTemp := vanilla.Map{
			"project_id": projectId,
		}
		tasks := b_task.NewTaskRepository(this.Ctx).GetByFilters(filtersTemp)
		respLogs := make([]*TaskOperationLog, 0)
		for _, task := range tasks {
			log := this.GetLogByTaskId(task.Id)
			respLogs = append(respLogs, log)
		}
		return respLogs
	} else {
		filters["task_id"] = taskId
		return this.GetByFilters(filters)
	}
}

func (this *TaskOperationLogRepository) GetLogByTaskId(taskId int) *TaskOperationLog{
	filters := vanilla.Map{
		"task_id": taskId,
	}
	logs := this.GetByFilters(filters)
	if len(logs) == 0 {
		return nil
	}
	return logs[0]
}

func NewTaskOperationLogRepository(ctx context.Context) *TaskOperationLogRepository {
	repository := new(TaskOperationLogRepository)
	repository.Ctx = ctx
	return repository
}