package account

import (
	"context"
	"github.com/kfchen81/beego"
	"github.com/kfchen81/beego/vanilla"
	m_account "teamdo/models/account"
)

type UserRepository struct {
	vanilla.RepositoryBase
}

func (this *UserRepository) GetByFilters(filters vanilla.Map) []*User {
	qs := vanilla.GetOrmFromContext(this.Ctx).QueryTable(&m_account.User{})
	if len(filters) > 0 {
		qs = qs.Filter(filters)
	}

	var dbModels []*m_account.User
	_, err := qs.OrderBy("-id").All(&dbModels)
	if err != nil {
		beego.Error(err)
		return []*User{}
	}
	Users := make([]*User, 0, len(dbModels))
	for _, dbModel := range dbModels {
		Users = append(Users, NewUserFromModel(this.Ctx, dbModel))
	}
	return Users
}

func NewUserRepository(ctx context.Context) *UserRepository {
	repository := new(UserRepository)
	repository.Ctx = ctx
	return repository
}
