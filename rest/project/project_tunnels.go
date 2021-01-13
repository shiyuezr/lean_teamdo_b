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

	b_project.NewFillTunnelsService(bCtx).Fill(tunnels, vanilla.FillOption{"with_options": true})
	rTunnels := b_project.NewEncodeTunnelService(bCtx).EncodeMany(tunnels)

	response := vanilla.MakeResponse(rTunnels)
	this.ReturnJSON(response)
}