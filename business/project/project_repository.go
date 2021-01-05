package project

import (
	"context"
	"github.com/kfchen81/beego"
	"github.com/kfchen81/beego/vanilla"
	m_project "lean_teamdo_b/models/project"
)

type ProjectRepository struct {
	vanilla.RepositoryBase
}

func (this *ProjectRepository) GetProjectByUserId(userId int) []*Project {
	filters := vanilla.Map{
		"user_id": userId,
	}
	projects := this.GetProject(filters)
	if len(projects) == 0 {
		return nil
	} else {
		return projects
	}
}

func (this *ProjectRepository) GetProject(filters vanilla.Map) []*Project {
	o := vanilla.GetOrmFromContext(this.Ctx)
	qs := o.QueryTable(&m_project.Project{})

	var models  []*m_project.Project
	if len(filters) > 0 {
		qs = qs.Filter(filters)
	}

	_, err := qs.All(&models)
	if err != nil {
		beego.Error(err)
		return nil
	}
	projects := make([]*Project, 0)
	for _, model := range models {
		projects = append(projects, NewProjectForModel(this.Ctx, model))
	}
	return projects
}

func NewProjectRepository(ctx context.Context) *ProjectRepository {
	repository := new(ProjectRepository)
	repository.Ctx = ctx
	return repository
}