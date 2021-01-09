package project

import (
	"github.com/kfchen81/beego/vanilla"
	b_project "teamdo/business/project"
)

type ProjectTunnel struct {
	vanilla.RestResource
}

func (this *ProjectTunnel) Resource() string {
	return "project.project_tunnel"
}

func (this *ProjectTunnel) GetParameters() map[string][]string {
	return map[string][]string{
		"PUT": []string{"project_id: int", "title: string"},
		"POST": []string{"id: int", "title: string"},
		"DELETE": []string{"id: int"},
	}
}

// 前端在创建项目的时候， 会判断当前用户是不是项目的管理员
func (this *ProjectTunnel) Put()  {
	bCtx := this.GetBusinessContext()

	projectId, _ := this.GetInt("project_id")
	title := this.GetString("title")
	b_project.NewTunnel(bCtx, projectId, title)
	response := vanilla.MakeResponse(vanilla.Map{})
	this.ReturnJSON(response)

}

func (this *ProjectTunnel) Post()  {
	bCtx := this.GetBusinessContext()
	id, _ := this.GetInt("id")
	title := this.GetString("title")

	tunnel := b_project.NewTunnelRepository(bCtx).GetTunnelById(id)
	tunnel.UpdateTitle(title)

	response := vanilla.MakeResponse(vanilla.Map{})
	this.ReturnJSON(response)
}

func (this *ProjectTunnel) Delete()  {
	bCtx := this.GetBusinessContext()
	id, _ := this.GetInt("id")

	tunnel := b_project.NewTunnelRepository(bCtx).GetTunnelById(id)
	tunnel.Deleted()
	response := vanilla.MakeResponse(vanilla.Map{})
	this.ReturnJSON(response)
}
