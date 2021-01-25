package project

import (
	"context"
	"github.com/kfchen81/beego/vanilla"
)

type EncodeProjectService struct {
	vanilla.ServiceBase
}

// Encode 对单个实体对象进行编码
func (this *EncodeProjectService) Encode(project *Project) *RProject {
	if project == nil {
		return nil
	}

	return &RProject{
		Id:          project.Id,
		Name:        project.Name,
		Information: project.Information,
	}
}

func NewEncodeProjectService(ctx context.Context) *EncodeProjectService {
	service := new(EncodeProjectService)
	service.Ctx = ctx
	return service
}
