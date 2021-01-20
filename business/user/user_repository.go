package user

import (
	"context"
	"github.com/kfchen81/beego"
	"github.com/kfchen81/beego/vanilla"
	m_user "teamdo/models/user"
)

type UserReository struct {
	vanilla.RepositoryBase
}

func (this *UserReository) GetByFilters(filters vanilla.Map) []*User {
	o := vanilla.GetOrmFromContext(this.Ctx)
	qs := o.QueryTable(&m_user.User{})
	if len(filters) != 0 {
		qs = qs.Filter(filters)
	}
	var models  []*m_user.User
	_, err := qs.All(&models)
	if err != nil {
		beego.Error(err)
		return nil
	}
	users := make([]*User, 0)
	for _, model := range models {
		users = append(users, NewUserForModel(this.Ctx, model))
	}
	return users
}

func NewUserRepository(ctx context.Context) *UserReository {
	repository := new(UserReository)
	repository.Ctx = ctx
	return repository
}
