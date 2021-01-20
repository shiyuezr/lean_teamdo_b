package tunnel

import (
	"github.com/kfchen81/beego/vanilla"
	b_tunnel "teamdo/business/tunnel"
)

type SortedTunnel struct {
	vanilla.RestResource
}

func (this *SortedTunnel) Resource() string {
	return "tunnel.sorted_tunnel"
}

func (this *SortedTunnel) GetParameters() map[string][]string {
	return map[string][]string{
		"PUT": []string{
			"ids:json-array",
		},
	}
}

func (this *SortedTunnel) Put()  {
	bCtx := this.GetBusinessContext()
	ids := this.GetIntArray("ids")

	if len(ids) > 1 {
		b_tunnel.NewTunnelRepository(bCtx).SortedTunnel(ids)
	}

	response := vanilla.MakeResponse(vanilla.Map{})
	this.ReturnJSON(response)
}
