package project

import (
	"context"
	"github.com/kfchen81/beego"
	"github.com/kfchen81/beego/orm"
	"github.com/kfchen81/beego/vanilla"
	m_project "teamdo/models/project"
)

type Task struct {
	vanilla.EntityBase
	Id 			int
	Title       string
	TunnelId	int
	ExecutorId  int
	projectName string

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

func (this *Task) UpdateStatus(status bool)  {
	o := vanilla.GetOrmFromContext(this.Ctx)
	
	_, err := o.QueryTable(&m_project.Task{}).Filter(vanilla.Map{"id": this.Id}).Update(orm.Params{
		"status": status,
	})

	if err != nil {
		beego.Error(err)
		panic(err)
	}
}

func (this *Task) UpdatePriority(priority string)  {
	o := vanilla.GetOrmFromContext(this.Ctx)

	_, err := o.QueryTable(&m_project.Task{}).Filter(vanilla.Map{"id": this.Id}).Update(orm.Params{
		"priority": priority,
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

func (this *Task) Update(title string, remark string)  {
	o := vanilla.GetOrmFromContext(this.Ctx)

	_, err := o.QueryTable(&m_project.Task{}).Filter(vanilla.Map{"id": this.Id}).Update(orm.Params{
		"title": title,
		"remark": remark,
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
	instance.Priority = dbModel.Priority
	instance.Remark = dbModel.Remark
	instance.projectName = dbModel.ProjectName
	return instance
}
