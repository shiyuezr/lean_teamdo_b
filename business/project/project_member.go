package project

import (
	"context"
	"teamdo/business/user"
	m_user "teamdo/models/user"
)

type ProjectMember struct {
	user.User
}

func NewProjectMemberForModel(ctx context.Context, dbModel *m_user.Member) *ProjectMember {
	instance := new(ProjectMember)
	instance.Ctx = ctx
	instance.Id = dbModel.Id
	instance.UserName = dbModel.UserName
	return instance
}
