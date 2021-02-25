package lane

import (
	"context"
	"fmt"
	"github.com/kfchen81/beego"
	"github.com/kfchen81/beego/orm"
	"github.com/kfchen81/beego/vanilla"
	"teamdo/business/project"
	m_lane "teamdo/models/lane"
)

type Lane struct {
	vanilla.EntityBase
	Id        int
	Name      string
	Sort      string
	ProjectId int

	Project *project.Project //所属项目
}

func (this *Lane) Update(name string, sort int) {

	var model m_lane.Lane
	o := vanilla.GetOrmFromContext(this.Ctx)
	if sort != 0 {
		_, err := o.QueryTable(&model).Filter("id", this.Id).Update(orm.Params{
			"sort": sort,
		})
		if err != nil {
			beego.Error(err)
			panic(vanilla.NewBusinessError("update：fail", fmt.Sprintf("修改项目失败")))
		}
		return
	}
	_, err := o.QueryTable(&model).Filter("id", this.Id).Update(orm.Params{
		"name": name,
	})
	if err != nil {
		beego.Error(err)
		panic(vanilla.NewBusinessError("update：fail", fmt.Sprintf("修改项目失败")))
	}
}

func (this *Lane) Delete() {
	_, err := vanilla.GetOrmFromContext(this.Ctx).QueryTable(m_lane.Lane{}).Filter(vanilla.Map{"id": this.Id}).Delete()
	if err != nil {
		beego.Error(err)
		panic(err)
	}
}

func NewLane(ctx context.Context, name string, pid int) *Lane {
	dbModel := &m_lane.Lane{Name: name, ProjectId: pid}
	_, err := vanilla.GetOrmFromContext(ctx).Insert(dbModel)

	if err != nil {
		beego.Error(err)
		panic(vanilla.NewSystemError("create_lane:failed", "创建泳道失败"))
	}
	return NewLaneFromDbModel(ctx, dbModel)
}

func NewLaneFromDbModel(ctx context.Context, dbModel *m_lane.Lane) *Lane {
	instance := new(Lane)
	instance.Ctx = ctx
	instance.Model = dbModel
	instance.Id = dbModel.Id
	instance.Name = dbModel.Name
	return instance
}
