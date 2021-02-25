package task_operation_log

import (
	"context"
	"github.com/kfchen81/beego/vanilla"
	b_account "teamdo/business/account"
	b_task "teamdo/business/task"
)

type FillTaskOperationLogService struct {
	vanilla.ServiceBase
}

func (this *FillTaskOperationLogService) FillOne(log *TaskOperationLog, option vanilla.FillOption) {
	this.Fill([]*TaskOperationLog{log}, option)
}

func (this *FillTaskOperationLogService) Fill(logs []*TaskOperationLog, option vanilla.FillOption) {
	if v, ok := option["with_task"]; ok && v {
		this.fillTask(logs)
	}
	if v, ok := option["with_operator"]; ok && v {
		this.fillOperator(logs)
	}
}

func (this *FillTaskOperationLogService) fillOperator(logs []*TaskOperationLog) {
	operatorIds := make([]int, 0)
	operatorId2operator := make(map[int]struct{})

	for _, log := range logs {
		if _, ok := operatorId2operator[log.UserId]; !ok {
			operatorIds = append(operatorIds, log.UserId)
		}
		operatorId2operator[log.UserId] = struct{}{}
	}
	operators := b_account.NewUserRepository(this.Ctx).GetByIds(operatorIds)
	id2operator := make(map[int]*b_account.User)
	for _, operator := range operators{
		id2operator[operator.Id] = operator
	}

	for _, log := range logs{
		log.User = id2operator[log.UserId]
	}
}

func (this *FillTaskOperationLogService) fillTask(logs []*TaskOperationLog) {
	taskIds := make([]int, 0)
	taskId2task := make(map[int]struct{})

	for _, log := range logs {
		if _, ok := taskId2task[log.TaskId]; !ok {
			taskIds = append(taskIds, log.TaskId)
		}
		taskId2task[log.TaskId] = struct{}{}
	}
	tasks := b_task.NewTaskRepository(this.Ctx).GetTaskByIds(taskIds)
	id2task := make(map[int]*b_task.Task)
	for _, task := range tasks{
		id2task[task.Id] = task
	}

	for _, log := range logs{
		log.Task = id2task[log.TaskId]
	}
}

func NewFillTaskOperationLogService(ctx context.Context) *FillTaskOperationLogService {
	inst := new(FillTaskOperationLogService)
	inst.Ctx = ctx
	return inst
}