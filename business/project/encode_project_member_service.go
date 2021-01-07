package project

import (
	"context"
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

	return &RProject{
		Name: project.Name,
	}
}

func (this *EncodeProjectService) EncodeMany(projects []*Project) []*RProject {
	rDatas := make([]*RProject, 0)
	for _, project := range projects {
		rDatas = append(rDatas, this.Encode(project))
	}
	return rDatas
}
