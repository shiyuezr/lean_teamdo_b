package user

import (
	"context"
	m_user "teamdo/models/user"
)

type ProjectMember struct {
	User
}

func NewProjectMemberForModel(ctx context.Context, dbModel *m_user.Member) *ProjectMember {
	instance := new(ProjectMember)
	instance.Ctx = ctx
	instance.Id = dbModel.Id
	instance.UserName = dbModel.UserName
	return instance
}
