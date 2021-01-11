package project

import (
	"context"
	"github.com/kfchen81/beego/vanilla"
)

type FillProjectService struct {
	vanilla.ServiceBase
}

func NewFillProjectService(ctx context.Context) *FillProjectService {
	service := new(FillProjectService)
	service.Ctx = ctx
	return service
}

func (this *FillProjectService) FillOne(project *Project, option vanilla.FillOption)  {
	if project == nil {
		return
	}

	if enableOption, ok := option["with_tunnel"]; ok && enableOption {
		this.FillTunnels(project)
	}
}


func (this *FillProjectService) FillTunnels(project *Project)  {

	tunnels := NewTunnelRepository(this.Ctx).GetTunnelsByProjectId(project.Id)
	if len(tunnels) == 0{
		return
	}
	tunnelIds := make([]int, 0)
	for _, tunnel := range tunnels {
		tunnelIds = append(tunnelIds, tunnel.Id)
	}
	this.FillTask(tunnels)

	project.Tunnel = tunnels
}

func (this *FillProjectService) FillTask(tunnels []*Tunnel)  {
	tunnelIds := make([]int, 0)
	for _, tunnel := range tunnels {
		tunnelIds = append(tunnelIds, tunnel.Id)
	}
	tasks := NewTaskRepository(this.Ctx).GetTasksByTunnelIds(tunnelIds)
	tunnelId2tasks :=make(map[int][]*Task)
	for _, task := range tasks {
		if exitTask, ok := tunnelId2tasks[task.Id]; ok {
			tunnelId2tasks[task.Id] = append(exitTask, task)
		} else {
			tunnelId2tasks[task.Id] = []*Task{task}
		}
	}

	for _, tunnel := range tunnels {
		tunnel.Task = tunnelId2tasks[tunnel.Id]
	}

}