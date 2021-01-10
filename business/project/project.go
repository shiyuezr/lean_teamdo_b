package project

import (
	"context"
	"fmt"
	"github.com/kfchen81/beego"
	"github.com/kfchen81/beego/vanilla"
	m_project "teamdo/models/project"
	"time"
)

type Project struct {
	vanilla.EntityBase
	Id				int
	Name    		string
	CreateAt		time.Time
	ManagerId		int

	Members			[]*Member
	Tunnel			[]*Tunnel
}

func NewProject(ctx context.Context, userId int, projectName string) *Project {
	o := vanilla.GetOrmFromContext(ctx)

	model := m_project.Project{}
	model.ManagerId = userId
	model.Name = projectName

	id, err := o.Insert(&model)
	if err != nil {
		beego.Error(err)
		panic(vanilla.NewBusinessError("create_project_fail", fmt.Sprintf("创建项目失败")))
	}
	model.Id = int(id)
	return NewProjectForModel(ctx, &model)
}


func NewProjectForModel(ctx context.Context, dbModel *m_project.Project) *Project {
	instance := new(Project)
	instance.Ctx = ctx
	instance.Id = dbModel.Id
	instance.Name = dbModel.Name
	instance.ManagerId = dbModel.ManagerId
	return instance
}