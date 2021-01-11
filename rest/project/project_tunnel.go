package project

import (
	"github.com/kfchen81/beego/vanilla"
	b_project "teamdo/business/project"
)

type ProjectTunnel struct {
	vanilla.RestResource
}

func (this *ProjectTunnel) Resource() string {
	return "project.tunnel"
}

func (this *ProjectTunnel) GetParameters() map[string][]string {
	return map[string][]string{
		"PUT": []string{"project_id:int", "title:string", "manager_id:int"},
		"POST": []string{"project_id:int", "id:int", "title:string", "manager_id:int"},
		"DELETE": []string{"project_id:int", "id:int", "manager_id:int"},
	}
}

// 前端在创建项目的时候， 会判断当前用户是不是项目的管理员
func (this *ProjectTunnel) Put()  {
	bCtx := this.GetBusinessContext()

	projectId, _ := this.GetInt("project_id")
	managerId, _ := this.GetInt("manager_id")
	title := this.GetString("title")

	project := b_project.NewProjectRepository(bCtx).GetProjectById(projectId)
	if project.ManagerId != managerId {
		panic(vanilla.NewBusinessError("not_project_manager","不是项目管理员"))
	}
	project.AddTunnel(title)
	response := vanilla.MakeResponse(vanilla.Map{})
	this.ReturnJSON(response)

}

func (this *ProjectTunnel) Post()  {
	bCtx := this.GetBusinessContext()

	id, _ := this.GetInt("id")
	title := this.GetString("title")
	managerId, _ := this.GetInt("manager_id")
	projectId, _ := this.GetInt("project_id")

	project := b_project.NewProjectRepository(bCtx).GetProjectById(projectId)
	if project.ManagerId != managerId {
		panic(vanilla.NewBusinessError("not_project_manager","不是项目管理员"))
	}
	tunnel := b_project.NewTunnelRepository(bCtx).GetTunnelById(id)
	tunnel.UpdateTitle(title)

	response := vanilla.MakeResponse(vanilla.Map{})
	this.ReturnJSON(response)
}

func (this *ProjectTunnel) Delete()  {
	bCtx := this.GetBusinessContext()

	id, _ := this.GetInt("id")
	managerId, _ := this.GetInt("manager_id")
	projectId, _ := this.GetInt("project_id")

	project := b_project.NewProjectRepository(bCtx).GetProjectById(projectId)
	if project.ManagerId != managerId {
		panic(vanilla.NewBusinessError("not_project_manager","不是项目管理员"))
	}
	tunnel := b_project.NewTunnelRepository(bCtx).GetTunnelById(id)
	tunnel.Deleted()
	response := vanilla.MakeResponse(vanilla.Map{})
	this.ReturnJSON(response)
}
