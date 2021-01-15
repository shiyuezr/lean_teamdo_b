package tunnel

import (
	"context"
	"github.com/kfchen81/beego/vanilla"
	b_project "teamdo/business/project"
	b_task "teamdo/business/task"
)

type EncodeTunnelService struct {
	vanilla.ServiceBase
}

func NewEncodeTunnelService(ctx context.Context) *EncodeTunnelService {
	service := new(EncodeTunnelService)
	service.Ctx = ctx
	return service
}

func (this *EncodeTunnelService) Encode(tunnel *Tunnel) *b_project.RTunnel {
	if tunnel == nil {
		return nil
	}

	rTasks := make([]*b_project.RTask, 0)
	if len(tunnel.Task) != 0 {
		for _, task := range tunnel.Task {
			rTasks = append(rTasks, b_task.NewEncodeTaskService(this.Ctx).Encode(task))
		}
	}

	return &b_project.RTunnel{
		Id: tunnel.Id,
		Title: tunnel.Title,
		Task: rTasks,
	}
}

func (this *EncodeTunnelService) EncodeMany(tunnels []*Tunnel) []*b_project.RTunnel {
	rDatas := make([]*b_project.RTunnel, 0)
	for _, tunnel := range tunnels {
		rDatas = append(rDatas, this.Encode(tunnel))
	}
	return rDatas
}