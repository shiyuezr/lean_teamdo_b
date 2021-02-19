package task

import (
	"context"
	"github.com/kfchen81/beego"
	"github.com/kfchen81/beego/vanilla"
	b_account "teamdo/business/account"
	m_account "teamdo/models/account"
	m_project "teamdo/models/project"
)

type FillTaskService struct {
	vanilla.ServiceBase
}

func (this *FillTaskService) FillOne(task *Task, option vanilla.FillOption) {
	this.Fill([]*Task{task}, option)
}

func (this *FillTaskService) Fill(tasks []*Task, option vanilla.FillOption) {
	ids := make([]int, 0)
	for _, task := range tasks {
		ids = append(ids, task.Id)
	}

	if v, ok := option["with_user"]; ok && v {
		this.fillUser(tasks, ids)
	}
	if v, ok := option["with_parent_task"]; ok && v {
		this.fillParticipant(tasks, ids)
	}
	if v, ok := option["with_lane"]; ok && v {
		this.fillParticipant(tasks, ids)
	}
	if v, ok := option["with_project"]; ok && v {
		this.fillParticipant(tasks, ids)
	}
}

func (this *FillTaskService) fillUser(tasks []*Task, ids []int) {
	id2entity := make(map[int]*Task)
	for _, task := range tasks {
		id2entity[task.Id] = task
	}

	corps := b_account.NewUserRepository(this.Ctx).GetById(corpIds)
	id2corp := make(map[int]*b_corp.Corp)
	for _, corp := range corps{
		id2corp[corp.Id] = corp
	}

	for _, store := range stores{
		store.Corp = id2corp[store.CorpId]
	}
}

func (this *FillTaskService) fillParticipant(projects []*Task, ids []int) {
	id2entity := make(map[int]*Task)
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
				project.Participants = append(project.Participants, b_account.NewUserFromModel(this.Ctx, model))
			}
		}
	}
}

func (this *FillTaskService) fillAdministrator(projects []*Task, ids []int) {
	id2entity := make(map[int]*Task)
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
				project.Administrators = append(project.Administrators, b_account.NewUserFromModel(this.Ctx, model))
			}
		}
	}
}

func NewFillTaskService(ctx context.Context) *FillTaskService {
	inst := new(FillTaskService)
	inst.Ctx = ctx
	return inst
}
