package tunnel

import (
	"github.com/kfchen81/beego/vanilla"
	_ "teamdo/business/account"
	"teamdo/business/tunnel"
)

type Tunnels struct {
	vanilla.RestResource
}

func (this *Tunnels) Resource() string {
	return "project.tunnels"
}

func (this *Tunnels) GetParameters() map[string][]string {
	return map[string][]string{
		"GET": []string{
			"project_id:int",
		},
	}
}

func (this *Tunnels) Get()  {
	bCtx := this.GetBusinessContext()
	projectId, _ := this.GetInt("project_id")

	tunnels := tunnel.NewTunnelRepository(bCtx).GetTunnelsByProjectId(projectId)

	tunnel.NewFillTunnelsService(bCtx).Fill(tunnels, vanilla.FillOption{"with_tasks": true})
	rTunnels := tunnel.NewEncodeTunnelService(bCtx).EncodeMany(tunnels)

	response := vanilla.MakeResponse(rTunnels)
	this.ReturnJSON(response)
}