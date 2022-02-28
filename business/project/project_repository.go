package project

import (
	"context"
	"github.com/kfchen81/beego"
	"github.com/kfchen81/beego/vanilla"
	models "teamdo/models/project"
	members "teamdo/models/project_member"
)

type ProjectRepository struct {
	vanilla.RepositoryBase
}

func (this *ProjectRepository) GetProjects(filters vanilla.Map, orderExprs ...string) []*Project {
	o := vanilla.GetOrmFromContext(this.Ctx)
	qs := o.QueryTable(&models.Project{})

	var models []*models.Project
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
		projects = append(projects, NewProjectFromDbModel(this.Ctx, model))
	}
	return projects
}



func (this *ProjectRepository) GetProject(id int) *Project {
	filters := vanilla.Map{
		"id": id,
	}

	projects := this.GetProjects(filters)

	if len(projects) == 0 {
		return nil
	} else {
		return projects[0]
	}
}

func (this *ProjectRepository)GetProjectByManagerId(managerId int)[]*Project{
	filters := vanilla.Map{
		"manager_id": managerId,
	}
	projects := this.GetProjects(filters)
	return projects
}

func (this *ProjectRepository)GetProjectByUserId(userId int)[]*Project{
	o := vanilla.GetOrmFromContext(this.Ctx)
	var projectMembers []*members.ProjectMember
	_,err := o.QueryTable(&members.ProjectMember{}).Filter("user_id",userId).All(&projectMembers)
	if err!=nil {
		beego.Error(err)
		return nil
	}
	projectIds:=make([]int,0)
	for k,_ := range projectMembers{
		projectIds = append(projectIds, projectMembers[k].ProjectId)
	}
	filters := vanilla.Map{
		"id__in": projectIds,
	}
	projects := this.GetProjects(filters)
	return projects
}

func NewProjectRepository(ctx context.Context) *ProjectRepository {
	inst := new(ProjectRepository)
	inst.Ctx = ctx
	return inst
}