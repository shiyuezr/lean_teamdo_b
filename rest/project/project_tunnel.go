package project

import (
	"github.com/kfchen81/beego/vanilla"
	b_project "lean_teamdo_b/business/project"
)

type ProjectTunnel struct {
	vanilla.RestResource
}

func (this *ProjectTunnel) Resource() string {
	return "project.project_tunnel"
}

func (this *ProjectTunnel) GetParameters() map[string][]string {
	return map[string][]string{
		"GET": []string{"project_id: int"},
		"PUT": []string{"project_id: int", "title: string", "manager_id: int"},
		"POST": []string{"id: int"},
		"DELETE": []string{"id: int"},
	}
}

func (this *ProjectTunnel) Get()  {

}

func (this *ProjectTunnel) Put()  {
	bCtx := this.GetBusinessContext()

	projectId, _ :=this.GetInt("project_id")
	title := this.GetString("title")
	managerId, _ := this.GetInt("manager_id")

	b_project.NewProjectTunnelFactory(bCtx).CreateProjectTunnel(projectId, managerId, title)

}

func (this *ProjectTunnel) Post()  {

}

func (this *ProjectTunnel) Delete()  {

}
