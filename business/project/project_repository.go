package project

import (
	"context"
	"github.com/kfchen81/beego"
	"github.com/kfchen81/beego/vanilla"
	m_project "teamdo/models/project"
)

type ProjectRepository struct {
	vanilla.RepositoryBase
}

func (this *ProjectRepository) GetByFilters(filters vanilla.Map) []*Project {
	qs := vanilla.GetOrmFromContext(this.Ctx).QueryTable(&m_project.Project{})
	if len(filters) > 0 {
		qs = qs.Filter(filters)
	}

	var dbModels []*m_project.Project
	_, err := qs.OrderBy("-id").All(&dbModels)
	if err != nil {
		beego.Error(err)
		return []*Project{}
	}
	Projects := make([]*Project, 0, len(dbModels))
	for _, dbModel := range dbModels {
		Projects = append(Projects, NewProjectFromDbModel(this.Ctx, dbModel))
	}
	return Projects
}

func (this *ProjectRepository) GetProjectById(id int) *Project {
	filters := vanilla.Map{
		"id": id,
	}
	projects := this.GetByFilters(filters)
	if len(projects) == 0 {
		return nil
	}
	return projects[0]
}

func NewProjectRepository(ctx context.Context) *ProjectRepository {
	repository := new(ProjectRepository)
	repository.Ctx = ctx
	return repository
}
