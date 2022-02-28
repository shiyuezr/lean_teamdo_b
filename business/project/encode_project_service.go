package project

import (
	"context"
	"github.com/kfchen81/beego/vanilla"
)

type EncodeProjectService struct {
	vanilla.ServiceBase
}

//Encode 对单个实体对象进行编码
func (this *EncodeProjectService) Encode(project *Project) *RProject {
	return &RProject{
		Id: project.Id,
		Name: project.Name,
		ManagerId: project.ManagerId,
		Detail: project.Detail,
	}
}

//EncodeMany 对实体对象进行批量编码
func (this *EncodeProjectService) EncodeMany(projects []*Project) []*RProject {
	rDatas := make([]*RProject, 0)
	for _, project := range projects {
		rDatas = append(rDatas, this.Encode(project))
	}
	return rDatas
}

func NewEncodeProjectService(ctx context.Context) *EncodeProjectService {
	service := new(EncodeProjectService)
	service.Ctx = ctx
	return service
}
