package tunnel

import (
	"context"
	"github.com/kfchen81/beego/vanilla"
	b_task "teamdo/business/task"
)

type FillTunnelsService struct {
	vanilla.ServiceBase
}

func NewFillTunnelsService(ctx context.Context) *FillTunnelsService {
	service := new(FillTunnelsService)
	service.Ctx = ctx
	return service
}

func (this *FillTunnelsService) Fill(tunnels []*Tunnel, option vanilla.FillOption)  {
	if len(tunnels) == 0 {
		return
	}

	tunnelIds := make([]int, 0)
	for _, tunnel := range tunnels {
		tunnelIds = append(tunnelIds, tunnel.Id)
	}

	if enableOption, ok := option["with_tunnel"]; ok && enableOption {
		this.FillTasks(tunnels)
	}
}

func (this *FillTunnelsService) FillTasks(tunnels []*Tunnel)  {
	tunnelIds := make([]int, 0)
	for _, tunnel := range tunnels {
		tunnelIds = append(tunnelIds, tunnel.Id)
	}
	tasks := b_task.NewTaskRepository(this.Ctx).GetTasksByTunnelIds(tunnelIds)
	tunnelId2tasks :=make(map[int][]*b_task.Task)
	for _, task := range tasks {
		if exitTask, ok := tunnelId2tasks[task.Id]; ok {
			tunnelId2tasks[task.Id] = append(exitTask, task)
		} else {
			tunnelId2tasks[task.Id] = []*b_task.Task{task}
		}
	}

	for _, tunnel := range tunnels {
		tunnel.Task = tunnelId2tasks[tunnel.Id]
	}
}
