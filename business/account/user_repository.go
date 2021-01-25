package account

import (
	"context"
	"github.com/kfchen81/beego/vanilla"
)

// todo 暂时当作数据库使用，后期切换mysql
var UserList = []*User{
	{
		Id:       1,
		Username: "tom",
		Password: "123456",
	},
	{
		Id:       2,
		Username: "jack",
		Password: "123456",
	},
}

type UserRepository struct {
	vanilla.RepositoryBase
}

func NewUserRepository(ctx context.Context) *UserRepository {
	repository := new(UserRepository)
	repository.Ctx = ctx
	return repository
}

func (this *UserRepository) GetUserByInformation(username string, password string) *User {
	resp_user := &User{}
	// todo 后期换成从数据库读取数据
	for _, value_user := range UserList {
		if username == value_user.Username && password == value_user.Password {
			resp_user = value_user
			return resp_user
		}
	}
	return resp_user
}

func (this *UserRepository) GetUsersById(id int) *User {
	resp_user := &User{}
	// todo 后期换成从数据库读取数据
	for _, value_user := range UserList {
		if id == value_user.Id {
			resp_user = value_user
			return resp_user
		}
	}
	return resp_user
}
