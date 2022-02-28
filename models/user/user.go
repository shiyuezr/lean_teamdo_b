package user

import (
	"github.com/kfchen81/beego/orm"
	"time"
)
type User struct {
	Id int
	UserName string
	Password string
	CreateTime time.Time
	IsDelete bool
}

func (this *User) TableName()  string {
	return "user_user"
}

func init()  {
	orm.RegisterModel(new(User))
}