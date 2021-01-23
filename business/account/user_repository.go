package account

import (
	"context"
	"github.com/kfchen81/beego/vanilla"
	//"teamdo/rest/account"
)

type UserRepository struct {
	vanilla.RepositoryBase
}

func NewUserRepository(ctx context.Context) *UserRepository {
	repository := new(UserRepository)
	repository.Ctx = ctx
	return repository
}

func (this *UserRepository) GetUser(id int) *User {
	u := User{
		Id:       1,
		Username: "to",
		Password: "12345",
		Token:    "111",
	}
	return &u
}
