package account

import (
	"time"

	"github.com/kfchen81/beego/orm"
)

type User struct {
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

func (self *User) TableName() string {
	return "account_user"
}

/*
func (this *User) TableIndex() [][]string {
	return [][]string{
		[]string{"IsEnabled", "IsDeleted"},
	}
}
*/
func init() {
	orm.RegisterModel(new(User))
}
