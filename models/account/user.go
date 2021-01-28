package account

import "github.com/kfchen81/beego/orm"

type User struct {
	Id       int `orm:"pk;auto"`
	Username string
	Password string
	Status   int `orm:"default(1)"`
}

func (self *User) TableName() string {
	return "account_user"
}

func init() {
	orm.RegisterModel(new(User))
}
