package project

import (
	"context"
	b_account "teamdo/business/account"

	"github.com/kfchen81/beego/vanilla"
)

type EncodeProjectService struct {
	vanilla.ServiceBase
}

func NewEncodeProjectService(ctx context.Context) *EncodeProjectService {
	service := new(EncodeProjectService)
	service.Ctx = ctx
	return service
}

func (this *EncodeProjectService) Encode(project *Project) *RProject {
	if project == nil {
		return nil
	}
	rUsers := b_account.NewEncodeUserService(this.Ctx).EncodeMany(project.Users)

	return &RProject{
		Id:           project.Id,
		Name:         project.Name,
		Introduction: project.Introduction,
		Cover:        project.Cover,
		Users:        rUsers,
		StartTime:    project.StartTime,
		CreateAt:     project.CreateAt,
		IsEnabled:    project.IsEnabled,
		IsDeleted:    project.IsDeleted,
	}
}
