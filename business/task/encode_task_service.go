package task

import (
	"context"
	"github.com/kfchen81/beego/vanilla"
	b_account "teamdo/business/account"
	b_lane "teamdo/business/lane"
	b_project "teamdo/business/project"
)

type EncodeTaskService struct {
	vanilla.ServiceBase
}

// Encode 对单个实体对象进行编码
func (this *EncodeTaskService) Encode(task *Task) *RTask {
	if task == nil {
		return nil
	}
	rOperator := b_account.NewEncodeUserService(this.Ctx).Encode(task.Operator)
	rParentTask := &RTask{}
	if task.ParentTask != nil {
		rParentTask = this.Encode(task.ParentTask)
	}
	rLane := b_lane.NewEncodeLaneService(this.Ctx).Encode(task.Lane)
	rProject := b_project.NewEncodeProjectService(this.Ctx).Encode(task.Project)

	encodedTask := &RTask{
		Id:      task.Id,
		Content: task.Content,
		Status:  task.Status,

		Operator:   rOperator,
		ParentTask: rParentTask,
		Lane:       rLane,
		Project:    rProject,
	}
	return encodedTask
}

func (this *EncodeTaskService) EncodeMany(tasks []*Task) []*RTask {
	encodedTasks := make([]*RTask, 0)
	for _, task := range tasks {
		encodedTasks = append(encodedTasks, this.Encode(task))
	}
	return encodedTasks
}

func NewEncodeTaskService(ctx context.Context) *EncodeTaskService {
	service := new(EncodeTaskService)
	service.Ctx = ctx
	return service
}
