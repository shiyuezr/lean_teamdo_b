package user

import (
	"github.com/kfchen81/beego/orm"
	"time"
)

type User struct {
	Id 			int
	UserName 	string
	Password 	string
}

func (self *User) TableName() string {
	return "user_user"
}

type Manager struct {
	Id 			int
	UserId 		int
	projectId	int
	CreatedAt	time.Time 	`orm:"auto_now_add;type(datetime)"`
}

func (self *Manager) TableName() string {
	return "user_manager"
}

type Executor struct {
	Id 			int
	UserId     	int
	TaskId 		int
	CreatedAt	time.Time 	`orm:"auto_now_add;type(datetime)"`
}

func (self *Executor) TableName() string {
	return "user_manager"
}

type Member struct {
	Id 			int
	UserName	string
	CreatedAt	time.Time 	`orm:"auto_now_add;type(datetime)"`
}

func (self *Member) TableName() string {
	return "user_member"
}

func init()  {
	orm.RegisterModel(new(User))
	orm.RegisterModel(new(Manager))
	orm.RegisterModel(new(Executor))
}