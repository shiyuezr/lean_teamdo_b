package project

import (
	"context"
	"fmt"
	"github.com/kfchen81/beego"
	"github.com/kfchen81/beego/vanilla"
	m_project "teamdo/models/project"
)

type ProjectFactory struct {
	vanilla.ServiceBase
}

func NewProjectFactory(ctx context.Context) *ProjectFactory {
	repository := new(ProjectFactory)
	repository.Ctx = ctx
	return repository
}

func (this *ProjectFactory) CreateProject(userId int, projectName string) *Project {
	o := vanilla.GetOrmFromContext(this.Ctx)

	model := m_project.Project{}
	model.ManagerId = userId
	model.Name = projectName

	_, err := o.Insert(&model)
	if err != nil {
		beego.Error(err)
		panic(vanilla.NewBusinessError("create_project_fail", fmt.Sprintf("创建项目失败")))
	}
	return NewProjectForModel(this.Ctx, &model)
}