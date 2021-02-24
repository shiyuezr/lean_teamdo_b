package user

import (
	"context"
	"github.com/kfchen81/beego/vanilla"
	m_user "teamdo/models/user"
)

type User struct {
	vanilla.EntityBase
	Id				int
	UserName 		string
	Password		string
}

func NewUserForModel(ctx context.Context, dbModel *m_user.User) *User {
	instance := new(User)
	instance.Ctx = ctx
	instance.UserName = dbModel.UserName
	instance.Password = dbModel.Password
	instance.Id = dbModel.Id
	return instance
}