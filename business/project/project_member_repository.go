package project

import (
	"context"
	"github.com/kfchen81/beego"
	"github.com/kfchen81/beego/vanilla"
	m_project "teamdo/models/project"
	m_user "teamdo/models/user"
)

type Member struct {
	vanilla.RepositoryBase
}

func NewMemberRepository(ctx context.Context) *Member {
	repository := new(Member)
	repository.Ctx = ctx
	return repository
}

func (this *Member) GetMembersByProjectId(projectId int) []*ProjectMember {
	filters := vanilla.Map{
		"project_id": projectId,
	}
	var models []m_project.ProjectHasMember
	o := vanilla.GetOrmFromContext(this.Ctx)
	_, err := o.QueryTable(m_project.ProjectHasMember{}).Filter(filters).All(&models)

	if err != nil {
		beego.Error(err)
		return nil
	}

	memberIds := make([]int, 0)
	for _, model := range models {
		memberIds = append(memberIds, model.Id)
	}

	if len(memberIds) == 0 {
		return nil
	}

	members := this.GetMembersByIds(memberIds)
	return members
}

func (this *Member) GetMembersByIds(ids []int) []*ProjectMember {
	filters := vanilla.Map{
		"id__in": ids,
	}
	members := this.GetByFilters(filters)
	if len(members) == 0 {
		return nil
	} else {
		return members
	}
}

func (this *Member) GetByFilters(filters vanilla.Map) []*ProjectMember {
	o := vanilla.GetOrmFromContext(this.Ctx)
	qs := o.QueryTable(&m_user.User{})

	var models []*m_user.User
	if len(filters) > 0 {
		qs = qs.Filter(filters)
	}
	_, err := qs.All(&models)
	if err != nil {
		beego.Error(err)
		return nil
	}
	members := make([]*ProjectMember, 0)
	for _, model := range models {
		members = append(members, NewProjectMemberForModel(this.Ctx, model))
	}

	return members
}

