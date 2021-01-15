package task

import (
	"context"
	"github.com/kfchen81/beego/vanilla"
	b_project "teamdo/business/project"
)

type EncodeTaskService struct {
	vanilla.ServiceBase
}

func NewEncodeTaskService(ctx context.Context) *EncodeTaskService {
	service := new(EncodeTaskService)
	service.Ctx = ctx
	return service
}

func (this *EncodeTaskService) Encode(task *Task) *b_project.RTask {
	if task == nil {
		return nil
	}

	return &b_project.RTask{
		Title: task.Title,
		Status: task.Status,
		Remark: task.Remark,
		Priority: task.Priority,
		StartDate: task.StartDate,
		EndDate: task.EndDate,
	}
}

func (this *EncodeTaskService) EncodeMany(tasks []*Task)  {
	rDatas := make([]*b_project.RTask, 0)
	for _, task := range tasks {
		rDatas = append(rDatas, this.Encode(task))
	}
}