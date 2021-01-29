package project

import (
	"context"
	"github.com/kfchen81/beego"
	"github.com/kfchen81/beego/vanilla"
	"teamdo/business/account"
	m_account "teamdo/models/account"
	m_project "teamdo/models/project"
)

type FillProjectService struct {
	vanilla.ServiceBase
}

func (this *FillProjectService) FillOne(project *Project, option vanilla.FillOption) {
	this.Fill([]*Project{project}, option)
}

func (this *FillProjectService) Fill(projects []*Project, option vanilla.FillOption) {
	ids := make([]int, 0)
	for _, project := range projects {
		ids = append(ids, project.Id)
	}

	if v, ok := option["with_administrators"]; ok && v {
		this.fillAdministrator(projects, ids)
	}
	if v, ok := option["with_participants"]; ok && v {
		this.fillParticipant(projects, ids)
	}
}

func (this *FillProjectService) fillParticipant(projects []*Project, ids []int) {
	id2entity := make(map[int]*Project)
	for _, project := range projects {
		id2entity[project.Id] = project
	}

	o := vanilla.GetOrmFromContext(this.Ctx)
	var relationModels []*m_project.ProjectToParticipants
	_, err := o.QueryTable(&m_project.ProjectToParticipants{}).Filter("project_id__in", ids).All(&relationModels)
	if err != nil {
		beego.Error(err)
		return
	}
	if len(relationModels) == 0 {
		return
	}

	participantIds := make([]int, 0)
	for _, relationModel := range relationModels {
		participantIds = append(participantIds, relationModel.ParticipantId)
	}
	var models []*m_account.User
	_, err = o.QueryTable(&m_account.User{}).Filter("id__in", participantIds).All(&models)
	if err != nil {
		beego.Error(err)
		return
	}
	id2model := make(map[int]*m_account.User)
	for _, model := range models {
		id2model[model.Id] = model
	}

	for _, relationModel := range relationModels {
		projectId := relationModel.ProjectId
		participantId := relationModel.ParticipantId

		if project, ok := id2entity[projectId]; ok {
			if model, ok2 := id2model[participantId]; ok2 {
				project.Participants = append(project.Participants, account.NewUserFromModel(this.Ctx, model))
			}
		}
	}
}

func (this *FillProjectService) fillAdministrator(projects []*Project, ids []int) {
	id2entity := make(map[int]*Project)
	for _, project := range projects {
		id2entity[project.Id] = project
	}
	o := vanilla.GetOrmFromContext(this.Ctx)
	var relationModels []*m_project.ProjectToAdministrators
	_, err := o.QueryTable(&m_project.ProjectToAdministrators{}).Filter("project_id__in", ids).All(&relationModels)
	if err != nil {
		beego.Error(err)
		return
	}
	if len(relationModels) == 0 {
		return
	}
	administratorIds := make([]int, 0)
	for _, relationModel := range relationModels {
		administratorIds = append(administratorIds, relationModel.AdministratorId)
	}
	var models []*m_account.User
	_, err = o.QueryTable(&m_account.User{}).Filter("id__in", administratorIds).All(&models)
	if err != nil {
		beego.Error(err)
		return
	}
	id2model := make(map[int]*m_account.User)
	for _, model := range models {
		id2model[model.Id] = model
	}
	for _, relationModel := range relationModels {
		projectId := relationModel.ProjectId
		administratorId := relationModel.AdministratorId

		if project, ok := id2entity[projectId]; ok {
			if model, ok2 := id2model[administratorId]; ok2 {
				project.Administrators = append(project.Administrators, account.NewUserFromModel(this.Ctx, model))
			}
		}
	}
}

func NewFillProjectService(ctx context.Context) *FillProjectService {
	inst := new(FillProjectService)
	inst.Ctx = ctx
	return inst
}
