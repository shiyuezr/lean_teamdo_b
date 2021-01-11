package project

import (
	"github.com/kfchen81/beego/vanilla"
	b_project "teamdo/business/project"
)

type ProjectTunnels struct {
	vanilla.RestResource
}

func (this *ProjectTunnels) Resource() string {
	return "project.tunnels"
}

func (this *ProjectTunnels) GetParameters() map[string][]string {
	return map[string][]string{
		"GET": []string{
			"project_id:int",
		},
	}
}

func (this *ProjectTunnels) Get()  {
	bCtx := this.GetBusinessContext()
	projectId, _ := this.GetInt("project_id")

	tunnels := b_project.NewTunnelRepository(bCtx).GetTunnelsByProjectId(projectId)
	data := b_project.NewEncodeTunnelService(bCtx).EncodeMany(tunnels)

	response := vanilla.MakeResponse(vanilla.Map{
		"tunnels": data,
	})
	this.ReturnJSON(response)
}