package project

import (
	"context"
	m_project "teamdo/models/project"

	"github.com/kfchen81/beego"
	"github.com/kfchen81/beego/vanilla"
)

type ProjectRepository struct {
	vanilla.RepositoryBase
}

//NewProjectRepository 返回仓储器
func NewProjectRepository(ctx context.Context) *ProjectRepository {
	repository := new(ProjectRepository)
	repository.Ctx = ctx
	return repository
}

//GetProjects 查询所有的Projects
func (this *ProjectRepository) GetProjects(filters vanilla.Map, orderExprs ...string) []*Project {
	o := vanilla.GetOrmFromContext(this.Ctx)

	qs := o.QueryTable(&m_project.Project{})

	var models []*m_project.Project
	//先过滤再排序
	if len(filters) > 0 {
		qs = qs.Filter(filters)
	}
	if len(orderExprs) > 0 {
		qs = qs.OrderBy(orderExprs...)
	}

	_, err := qs.All(&models)

	if err != nil {
		beego.Error(err)
		return nil
	}

	projects := make([]*Project, 0)

	for _, model := range models {
		projects = append(projects, NewProjectFromModel(this.Ctx, model))
	}
	return projects
}

//GetProject 使用id查询Project
func (this *ProjectRepository) GetProject(id int) *Project {
	filter := vanilla.Map{
		"id": id,
	}

	/*****以下操作实现了GetSingleOrDefault****/
	products := this.GetProjects(filter)

	if len(products) <= 0 {
		return nil
	}

	return products[0]
	/**************************************/
}
