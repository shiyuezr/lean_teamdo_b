package user

import (
	"context"
	"github.com/kfchen81/beego"
	"github.com/kfchen81/beego/vanilla"
	models "teamdo/models/user"
)
type UserRepository struct {
	vanilla.RepositoryBase
}

func (this *UserRepository) GetUsers(filters vanilla.Map, orderExprs ...string) []*User {
	o := vanilla.GetOrmFromContext(this.Ctx)
	qs := o.QueryTable(&models.User{})

	var models []*models.User
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

	users := make([]*User, 0)
	for _, model := range models {
		users = append(users, NewUserFromDbModel(this.Ctx, model))
	}
	return users
}



func (this *UserRepository) GetUser(id int) *User {
	filters := vanilla.Map{
		"id": id,
	}

	projects := this.GetUsers(filters)

	if len(projects) == 0 {
		return nil
	} else {
		return projects[0]
	}
}

func (this *UserRepository) GetUserByUserNameAndPassword(userName string,password string) *User {
	filters := vanilla.Map{
		"user_name": userName,
		"password":password,
	}

	user := this.GetUsers(filters)

	if len(user) == 0 {
		return nil
	} else {
		return user[0]
	}
}

func NewUserRepository(ctx context.Context) *UserRepository {
	inst := new(UserRepository)
	inst.Ctx = ctx
	return inst
}