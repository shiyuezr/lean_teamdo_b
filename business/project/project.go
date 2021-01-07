package project

import (
	"context"
	"fmt"
	"github.com/kfchen81/beego"
	"github.com/kfchen81/beego/vanilla"
	"teamdo/business/user"
	m_project "teamdo/models/project"
	"time"
)

type Project struct {
	vanilla.EntityBase
	Id				int
	Name    		string
	CreateAt		time.Time
	ManagerId		int

	Members			[]*user.Member
	Tunnel			[]*Tunnel
}


func NewProjectForModel(ctx context.Context, dbModel *m_project.Project) *Project {
	instance := new(Project)
	instance.Ctx = ctx
	instance.Id = dbModel.Id
	instance.Name = dbModel.Name
	instance.ManagerId = dbModel.ManagerId
	return instance
}

func CreateProject(ctx context.Context, userId int, projectName string) *Project {
	o := vanilla.GetOrmFromContext(ctx)

	model := m_project.Project{}
	model.ManagerId = userId
	model.Name = projectName

	_, err := o.Insert(&model)
	if err != nil {
		beego.Error(err)
		panic(vanilla.NewBusinessError("create_project_fail", fmt.Sprintf("创建项目失败")))
	}
	return NewProjectForModel(ctx, &model)
}