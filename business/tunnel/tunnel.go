package tunnel

import (
	"context"
	"fmt"
	"github.com/kfchen81/beego"
	"github.com/kfchen81/beego/orm"
	"github.com/kfchen81/beego/vanilla"
	"strings"
	"teamdo/business/constant"
	b_task "teamdo/business/task"
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
	DisplayIndex int

	Tasks        []*b_task.Task
}

type TaskParams struct {
	ExecutorId		int
	Title			string
	Remark			string
	Priority 		string
	StartDate		time.Time
	EndDate			time.Time
}

func (this *Tunnel) Sorted(action string)  {
	displayIndex := 0
	if action == "left" {
		displayIndex = -1
	} else {
		displayIndex = 1
	}
	o := vanilla.GetOrmFromContext(this.Ctx)
	_, err := o.QueryTable(&m_project.Tunnel{}).Filter("id", this.Id).Update(orm.Params{
		"display_index": this.DisplayIndex + displayIndex,
	})

	if err != nil {
		beego.Error(err)
		panic(err)
	}
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

func (this *Tunnel) AddTask(taskParams *TaskParams)  {
	o := vanilla.GetOrmFromContext(this.Ctx)

	model := m_project.Task{}
	model.TunnelId = this.Id
	model.Title = taskParams.Title
	model.Remark = taskParams.Remark
	model.Priority = m_project.TASK_PRIOTITY_TYPE_CODE2PRIOTITY_TYPE[taskParams.Priority]
	model.ExecutorId = taskParams.ExecutorId
	model.StartDate = taskParams.StartDate
	model.EndDate = taskParams.EndDate

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
	instance.DisplayIndex = dbModel.DisplayIndex
	return instance
}

func NewTaskParams(
	executorId int,
	title string,
	remark string,
	priority string,
	startDateStr string,
	endDateStr 	string,
	) *TaskParams {
	startDate, _ := time.ParseInLocation(
		constant.TIME_LAYOUT,
		strings.ReplaceAll(startDateStr, "/", "-"),
		time.Local)
	endDate, _ := time.ParseInLocation(
		constant.TIME_LAYOUT,
		strings.ReplaceAll(endDateStr, "/", "-"),
		time.Local)
	beego.Info("7878787&*&*",startDate, "*&*&*")
	return &TaskParams{
		ExecutorId: executorId,
		Title: title,
		Remark: remark,
		Priority: priority,
		StartDate: startDate,
		EndDate: endDate,
	}
}




