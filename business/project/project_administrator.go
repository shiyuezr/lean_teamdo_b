package project

import (
	"context"
	"github.com/kfchen81/beego"
	"github.com/kfchen81/beego/vanilla"
	m_project "teamdo/models/project"
)

type ProjectToAdministrators struct {
	vanilla.EntityBase
	Id               int
	Project_id       int
	Administrator_id int
}

func NewProjectToAdministrator(ctx context.Context, pid int, uid int) *ProjectToAdministrators {
	o := vanilla.GetOrmFromContext(ctx)

	dbModel := &m_project.ProjectToAdministrators{
		ProjectId:    pid,
		AdministratorId: uid,
	}
	_, err := o.Insert(dbModel)
	if err != nil {
		beego.Error(err)
		panic(vanilla.NewSystemError("create:failed", "创建失败"))
	}
	return NewProjectToAdministratorFromDbModel(ctx, dbModel)
}

func NewProjectToAdministratorFromDbModel(ctx context.Context, dbModel *m_project.ProjectToAdministrators) *ProjectToAdministrators {
	instance := new(ProjectToAdministrators)
	instance.Ctx = ctx
	instance.Model = dbModel
	instance.Id = dbModel.Id
	instance.Project_id = dbModel.ProjectId
	instance.Administrator_id = dbModel.AdministratorId
	return instance
}


