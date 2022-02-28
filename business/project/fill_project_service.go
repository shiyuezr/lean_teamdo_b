package project

import (
	"context"
	"github.com/kfchen81/beego/vanilla"
)


type FillProjectService struct {
	vanilla.ServiceBase
}


func NewFillProjectService(ctx context.Context) *FillProjectService {
	service := new(FillProjectService)
	service.Ctx = ctx
	return service
}

func (this *FillProjectService) FillOne(project *Project, option vanilla.FillOption) {
	this.Fill([]*Project{ project }, option)
}

func (this *FillProjectService) Fill(projects []*Project, option vanilla.FillOption) {
	if len(projects) == 0 {
		return
	}

	ids := make([]int, 0)
	for _, project := range projects {
		ids = append(ids, project.Id)
	}

	if enableOption, ok := option["with_project_member"]; ok && enableOption {
		this.fillProjectMember(projects, ids)
	}


	return
}

func (this *FillProjectService) fillProjectMember(projects []*Project, ids []int) {
	////获取关联的id集合
	//ProjectManagerIds := make([]int, 0)
	//for _, project := range projects {
	//	ProjectManagerIds = append(ProjectManagerIds, project.ManagerId)
	//}
	//
	//
	//var projectMembers []*m_models.ProjectMember
	//o := vanilla.GetOrmFromContext(this.Ctx)
	//_, err := o.QueryTable(&m_models.ProjectMember{}).Filter("UserId__in", ProjectManagerIds).All(&projectMembers)
	//if err != nil {
	//	beego.Error(err)
	//	return
	//}
	//
	//projectId2models := make(map[int] []*m_models.ProjectMember )
	//for _, projectMember := range projectMembers {
	//	projectId2models[projectMember.UserId] = append(projectId2models[projectMember.UserId],projectMember)
	//}
	//
	//for _, project := range projects {
	//	project.ProjectMembers=project_member.NewProjectMembersFromDbModels(this.Ctx,projectId2models[project.Id])
	//}
}