package project

import (
	"context"
	"github.com/kfchen81/beego/vanilla"
)

type EncodeTaskService struct {
	vanilla.ServiceBase
}

func NewEncodeTaskService(ctx context.Context) *EncodeTaskService {
	service := new(EncodeTaskService)
	service.Ctx = ctx
	return service
}

func (this *EncodeTaskService) Encode(task *Task) *RTask {
	if task == nil {
		return nil
	}

	return &RTask{
		Title: task.Title,
		ProjectName: task.projectName,
		Status: task.Status,
		Remark: task.Remark,
		Priority: task.Priority,
		StartDate: task.StartDate,
		EndDate: task.EndDate,
	}
}

func (this *EncodeTaskService) EncodeMany(tasks []*Task)  {
	rDatas := make([]*RTask, 0)
	for _, task := range tasks {
		rDatas = append(rDatas, this.Encode(task))
	}
}