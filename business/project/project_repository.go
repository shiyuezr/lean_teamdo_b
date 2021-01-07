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

func (this *ProjectRepository) GetPagedProjects(filters vanilla.Map, page *vanilla.PageInfo, orderExprs ...string) ([]*Project, vanilla.INextPageInfo) {
	o := vanilla.GetOrmFromContext(this.Ctx)
	qs := o.QueryTable(&m_project.Project{})

	var models []*m_project.Project
	if len(filters) > 0 {
		qs = qs.Filter(filters)
	}
	if len(orderExprs) > 0 {
		qs = qs.OrderBy(orderExprs...)
	}

	paginateResult, err := vanilla.Paginate(qs, page, &models)

	if err != nil {
		beego.Error(err)
		return nil, paginateResult
	}

	projects := make([]*Project, 0)
	for _, model := range models {
		projects = append(projects, NewProjectForModel(this.Ctx, model))
	}

	return projects, paginateResult
}

func (this *ProjectRepository) GetProjectsByUserId(userId int) []*Project {
	filters := vanilla.Map{
		"user_id": userId,
	}
	var models []*m_project.ProjectHasUser
	o := vanilla.GetOrmFromContext(this.Ctx)
	_, err := o.QueryTable(&m_project.ProjectHasUser{}).Filter(filters).All(&models)
	if err != nil {
		beego.Error(err)
		return nil
	}

	ids := make([]int, 0)
	for _, model := range models {
		ids = append(ids, model.ProjectId)
	}

	projects := this.GetProjectsByIds(ids)
	return projects
}

func (this *ProjectRepository) GetProjectsByIds(ids []int) []*Project {
	filters := vanilla.Map{
		"id__in": ids,
	}
	projects := this.GetByFilters(filters)
	if len(projects) == 0 {
		return nil
	} else {
		return projects
	}
}

func (this *ProjectRepository) GetByFilters(filters vanilla.Map) []*Project {
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