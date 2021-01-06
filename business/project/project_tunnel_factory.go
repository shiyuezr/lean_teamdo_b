package project

import (
	"context"
	"fmt"
	"github.com/kfchen81/beego"
	"github.com/kfchen81/beego/vanilla"
	m_project "teamdo/models/project"
)

type ProjectTunnelFactory struct {
	vanilla.ServiceBase
}

func NewProjectTunnelFactory(ctx context.Context) *ProjectTunnelFactory {
	repository := new(ProjectTunnelFactory)
	repository.Ctx = ctx
	return repository
}

func (this *ProjectTunnelFactory) CreateProjectTunnel(projectId int, managerId int ,title string)  {
	o := vanilla.GetOrmFromContext(this.Ctx)

	model := m_project.Tunnel{}
	model.ProjectId = projectId
	model.Title = title
	model.ManagerId = managerId

	_, err := o.Insert(&model)
	if err != nil {
		beego.Error(err)
		panic(vanilla.NewBusinessError("create_project_tunnel_fail", fmt.Sprintf("创建泳道失败")))
	}
}