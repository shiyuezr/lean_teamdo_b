package project

import (
	"context"
	"fmt"
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

func (this *Task) UpdateExecutor(executorId int)  {
	
}

func (this *Task) Update(title string, status bool, remark string, priority string)  {
	o := vanilla.GetOrmFromContext(this.Ctx)

	_, err := o.QueryTable(&m_project.Task{}).Filter(vanilla.Map{"id": this.Id}).Update(orm.Params{
		"title": title,
		"remark": remark,
		"priority": priority,
	})

	if err != nil {
		beego.Error(err)
		panic(err)
	}
}

func NewTask(
	ctx context.Context,
	title string,
	tunnelId int,
	executorId int,
	status 	bool,
	remark	string,
	priority string,
	projectName string,
	) *Task {

	o := vanilla.GetOrmFromContext(ctx)

	model := m_project.Task{}
	model.Title = title
	model.ExecutorId = executorId
	model.TunnelId = tunnelId
	model.Status = status
	model.Remark = remark
	model.ProjectName = projectName
	model.Priority = priority

	id, err := o.Insert(&model)
	if err != nil {
		beego.Error(err)
		panic(vanilla.NewBusinessError("create_task_fail", fmt.Sprintf("创建任务失败")))
	}
	model.Id = int(id)

	return NewTaskForModel(ctx, &model)

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
