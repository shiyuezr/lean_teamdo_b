package user

import "github.com/kfchen81/beego/orm"

type User struct {
	Id 			int
	UserName 	string
}

func (self *User) TableName() string {
	return "user_user"
}

type Manager struct {
	Id 			int
	UserName	string
}

func (self *Manager) TableName() string {
	return "user_manager"
}

type Executor struct {
	Id 			int
	UserName	string
}

func (self *Executor) TableName() string {
	return "user_manager"
}

type Member struct {
	Id 			int
	UserName	string
}

func (self *Member) TableName() string {
	return "user_member"
}

func init()  {
	orm.RegisterModel(new(User))
	orm.RegisterModel(new(Manager))
	orm.RegisterModel(new(Executor))
}