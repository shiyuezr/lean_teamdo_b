package project

import (
	"context"
	"fmt"
	"github.com/kfchen81/beego"
	"github.com/kfchen81/beego/orm"
	"github.com/kfchen81/beego/vanilla"
	"teamdo/business/constant"
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

func (this *Tunnel) AddTask(executorId int, title, remark, priority, startDateStr, endDateStr string)  {
	startDate, _ := time.ParseInLocation(constant.DATE_LAYOUT, startDateStr, time.Local)
	endDate, _ := time.ParseInLocation(constant.DATE_LAYOUT, endDateStr, time.Local)

	o := vanilla.GetOrmFromContext(this.Ctx)
	model := m_project.Task{}
	model.Title = title
	model.Remark = remark
	model.Priority = priority
	model.ExecutorId = executorId
	model.StartDate = startDate
	model.EndDate = endDate

	_, err := o.Insert(&model)
	if err != nil {
		beego.Error(err)
		panic(vanilla.NewBusinessError("create_task_fail",fmt.Sprintf("添加任务失败")))
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
	instance.Id = dbModel.Id
	instance.ProjectId = dbModel.ProjectId
	instance.Title = dbModel.Title
	return instance
}




