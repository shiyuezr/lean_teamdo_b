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
	}

	rTasks := make([]*RTask, 0)
	if len(tunnel.Task) != 0 {
		for _, task := range tunnel.Task {
			rTasks = append(rTasks, NewEncodeTaskService(this.Ctx).Encode(task))
		}
	}

	return &RTunnel{
		Id: tunnel.Id,
		Title: tunnel.Title,
		Task: rTasks,
	}
}

func (this *EncodeTunnelService) EncodeMany(tunnels []*Tunnel) []*RTunnel {
	rDatas := make([]*RTunnel, 0)
	for _, tunnel := range tunnels {
		rDatas = append(rDatas, this.Encode(tunnel))
	}
	return rDatas
}