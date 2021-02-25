package task_operation_log

import (
	"context"
	"github.com/kfchen81/beego/vanilla"
	b_account "teamdo/business/account"
	b_task "teamdo/business/task"
)

type EncodeTaskOperationLogService struct {
	vanilla.ServiceBase
}

// Encode 对单个实体对象进行编码
func (this *EncodeTaskOperationLogService) Encode(log *TaskOperationLog) *RTaskOperationLog {
	if log == nil {
		return nil
	}
	rOperator := b_account.NewEncodeUserService(this.Ctx).Encode(log.User)
	rTask := b_task.NewEncodeTaskService(this.Ctx).Encode(log.Task)

	encodedLog := &RTaskOperationLog{
		Id:   log.Id,
		Content: log.Content,
		CreatedAt: log.CreatedAt.Format("2006-01-02 15:04:05"),
		Operator: rOperator,
		Task: rTask,
	}
	return encodedLog
}

func (this *EncodeTaskOperationLogService) EncodeMany(logs []*TaskOperationLog) []*RTaskOperationLog {
	encodedLogs := make([]*RTaskOperationLog, 0)
	for _, log := range logs {
		encodedLogs = append(encodedLogs, this.Encode(log))
	}
	return encodedLogs
}

func NewEncodeTaskOperationLogService(ctx context.Context) *EncodeTaskOperationLogService {
	service := new(EncodeTaskOperationLogService)
	service.Ctx = ctx
	return service
}