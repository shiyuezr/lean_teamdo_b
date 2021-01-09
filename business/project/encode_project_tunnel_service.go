package project

import (
	"context"
	"github.com/kfchen81/beego/vanilla"
)

type EncodeTunnelService struct {
	vanilla.ServiceBase
}

func NewEncodeTunnelService(ctx context.Context) *EncodeTunnelService {
	service := new(EncodeTunnelService)
	service.Ctx = ctx
	return service
}

func (this *EncodeTunnelService) Encode(tunnel *Tunnel) *RTunnel {
	if tunnel == nil {
		return nil
	} else {
		return &RTunnel{
			Id: tunnel.Id,
			Title: tunnel.Title,
		}
	}


}

func (this *EncodeTunnelService) EncodeMany(tunnels []*Tunnel) []*RTunnel {
	rDatas := make([]*RTunnel, 0)
	for _, tunnel := range tunnels {
		rDatas = append(rDatas, this.Encode(tunnel))
	}
	return rDatas
}