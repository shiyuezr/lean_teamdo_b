package task

import (
	"context"
	"github.com/kfchen81/beego"
	"github.com/kfchen81/beego/orm"
	"github.com/kfchen81/beego/vanilla"
	"strings"
	m_project "teamdo/models/project"
)

type Task struct {
	vanilla.EntityBase
	Id 			int
	Title       string
	TunnelId	int
	ExecutorId  int

	Status 		bool
	IsDelete	bool
	Remark		string
	Priority	string
	StartDate	string
	EndDate		string
}

func (this *Task) Delete()  {
	o := vanilla.GetOrmFromContext(this.Ctx)
	_, err := o.QueryTable(&m_project.Task{}).Filter(vanilla.Map{"id": this.Id}).Delete()

	if err != nil {
		beego.Error(err)
		panic(err)
	}
}

func (this *Task) UpdateTitle(title string)  {
	o := vanilla.GetOrmFromContext(this.Ctx)

	_, err := o.QueryTable(&m_project.Task{}).Filter(vanilla.Map{"id": this.Id}).Update(orm.Params{
		"title": title,
	})

	if err != nil {
		beego.Error(err)
		panic(err)
	}
}

func (this *Task) FinishTask()  {
	o := vanilla.GetOrmFromContext(this.Ctx)
	
	_, err := o.QueryTable(&m_project.Task{}).Filter(vanilla.Map{"id": this.Id}).Update(orm.Params{
		"status": true,
	})

	if err != nil {
		beego.Error(err)
		panic(err)
	}
}

func (this *Task) UnFinishTask()  {
	o := vanilla.GetOrmFromContext(this.Ctx)

	_, err := o.QueryTable(&m_project.Task{}).Filter(vanilla.Map{"id": this.Id}).Update(orm.Params{
		"status": false,
	})

	if err != nil {
		beego.Error(err)
		panic(err)
	}
}

func (this *Task) UpdatePriority(priority string)  {
	o := vanilla.GetOrmFromContext(this.Ctx)

	_, err := o.QueryTable(&m_project.Task{}).Filter(vanilla.Map{"id": this.Id}).Update(orm.Params{
		"priority": m_project.TASK_PRIOTITY_TYPE_CODE2PRIOTITY_TYPE[priority],
	})

	if err != nil {
		beego.Error(err)
		panic(err)
	}
}


func (this *Task) UpdateExecutor(userId int)  {
	o := vanilla.GetOrmFromContext(this.Ctx)

	_, err := o.QueryTable(&m_project.Task{}).Filter(vanilla.Map{"id": this.Id}).Update(orm.Params{
		"executor_id": userId,
	})

	if err != nil {
		beego.Error(err)
		panic(err)
	}
}

func (this *Task) Update(remark string, title string)  {
	o := vanilla.GetOrmFromContext(this.Ctx)

	_, err := o.QueryTable(&m_project.Task{}).Filter(vanilla.Map{"id": this.Id}).Update(orm.Params{
		"remark": remark,
		"title": title,
	})

	if err != nil {
		beego.Error(err)
		panic(err)
	}
}

func NewTaskForModel(ctx context.Context, dbModel *m_project.Task) *Task {
	instance := new(Task)
	instance.Ctx = ctx
	instance.Id = dbModel.Id
	instance.Title = dbModel.Title
	instance.Status = dbModel.Status
	instance.Priority = m_project.PRIOTITY_TYPE2CODE[dbModel.Priority]
	instance.Remark = dbModel.Remark
	instance.TunnelId = dbModel.TunnelId
	instance.ExecutorId = dbModel.ExecutorId
	instance.StartDate = strings.ReplaceAll(dbModel.StartDate.Format("2006-01-02 15:04:05"), "-", "/")
	instance.EndDate = strings.ReplaceAll(dbModel.EndDate.Format("2006-01-02 15:04:05"), "-", "/")
	return instance
}
