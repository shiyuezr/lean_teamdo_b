package project

import (
	"context"
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




