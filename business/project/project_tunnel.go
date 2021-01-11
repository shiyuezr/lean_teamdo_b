package project

import (
	"context"
	"fmt"
	"github.com/kfchen81/beego"
	"github.com/kfchen81/beego/orm"
	"github.com/kfchen81/beego/vanilla"
	m_project "teamdo/models/project"
	"time"
)

type Tunnel struct {
	vanilla.EntityBase
	Id			int
	Title 		string
	ProjectId   int
	IsDelete    bool
	CreateAt    time.Time

	Task        []*Task
}

func (this *Tunnel) UpdateTitle(title string)  {
	o := vanilla.GetOrmFromContext(this.Ctx)
	_, err := o.QueryTable(&m_project.Tunnel{}).Filter("id", this.Id).Update(orm.Params{
		"title": title,
	})
	if err != nil {
		beego.Error(err)
		panic(err)
	}
}

func (this *Tunnel) Deleted()  {
	o := vanilla.GetOrmFromContext(this.Ctx)
	_, err := o.QueryTable(&m_project.Tunnel{}).Filter("id", this.Id).Delete()

	if err != nil {
		beego.Error(err)
		panic(err)
	}
}
func NewTunnelForModel(ctx context.Context, dbModel *m_project.Tunnel) *Tunnel {
	instance := new(Tunnel)
	instance.Ctx = ctx
	instance.ProjectId = dbModel.ProjectId
	instance.Title = dbModel.Title
	return instance
}

//工厂方法
func NewTunnel(ctx context.Context,projectId int, title string) *Tunnel {
	o := vanilla.GetOrmFromContext(ctx)
	model := m_project.Tunnel{}
	model.Title = title
	model.ProjectId = projectId

	id, err := o.Insert(&model)
	if err != nil {
		beego.Error(err)
		panic(vanilla.NewBusinessError("tunnel_create_fail",fmt.Sprintf("创建泳道失败")))
	}
	model.Id = int(id)

	return NewTunnelForModel(ctx, &model)
}


