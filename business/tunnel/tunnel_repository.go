package tunnel

import (
	"context"
	"github.com/kfchen81/beego"
	"github.com/kfchen81/beego/orm"
	"github.com/kfchen81/beego/vanilla"
	m_project "teamdo/models/project"
)

type TunnelRepository struct {
	vanilla.ServiceBase
}

func NewTunnelRepository(ctx context.Context) *TunnelRepository {
	repository := new(TunnelRepository)
	repository.Ctx = ctx
	return repository
}

func (this *TunnelRepository) SortTunnels(ids []int)  {
	index := 0
	for _, id := range ids {
		this.UpdateDisplayIndex(id, index)
		index += 1
	}
}

func (this *TunnelRepository) UpdateDisplayIndex(id int, index int)  {
	o := vanilla.GetOrmFromContext(this.Ctx)
	_, err := o.QueryTable(&m_project.Tunnel{}).Filter(vanilla.Map{"id": id}).Update(orm.Params{
		"display_index": index,
	})

	if err != nil {
		beego.Error(err)
		panic(err)
	}
}

func (this *TunnelRepository) GetByFilters(filters vanilla.Map, orderExprs ...string) []*Tunnel {
	o := vanilla.GetOrmFromContext(this.Ctx)
	qs := o.QueryTable(&m_project.Tunnel{})

	if len(filters) > 0 {
		qs = qs.Filter(filters)
	}

	if len(orderExprs) > 0 {
		qs = qs.OrderBy(orderExprs...)
	} else {
		qs = qs.OrderBy("display_index")
	}
	var models []*m_project.Tunnel
	_, err := qs.All(&models)
	if err != nil {
		beego.Error(err)
		return nil
	}
	tunnels := make([]*Tunnel, 0)
	for _, model := range models {
		tunnels = append(tunnels, NewTunnelForModel(this.Ctx,model))
	}
	return tunnels
}

func (this *TunnelRepository) GetTunnelsByProjectId(projectId int) []*Tunnel {
	filters := vanilla.Map{
		"project_id": projectId,
		"is_deleted": false,
	}

	tunnels := this.GetByFilters(filters)
	if len(tunnels) == 0 {
		return nil
	} else {
		return tunnels
	}
}

func (this *TunnelRepository) GetTunnelById(id int) *Tunnel {
	filters := vanilla.Map{
		"id": id,
		"is_deleted": false,
	}
	tunnels := this.GetByFilters(filters)
	if len(tunnels) >0 {
		return tunnels[0]
	} else {
		return nil
	}
}
