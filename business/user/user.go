package user

import "github.com/kfchen81/beego/vanilla"

type User struct {
	vanilla.EntityBase
	Id				int
	UserName 		string
	Password		string
}