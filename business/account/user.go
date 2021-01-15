package account

import (
	"context"
	m_account "teamdo/models/account"
	"time"

	"github.com/kfchen81/beego/vanilla"
)

type User struct {
	vanilla.EntityBase
	Id        int
	UserName  string
	UserCode  string
	Age       int
	Password  string
	Level     int //0为管理员，1为普通成员
	IsEnabled bool
	IsDeleted bool
	CreateAt  time.Time
}

func NewUserFromModel(ctx context.Context, model *m_account.User) *User {
	instance := new(User)
	instance.Id = model.Id
	instance.UserCode = model.UserCode
	instance.UserName = model.UserName

	return instance
}
func init() {

}
