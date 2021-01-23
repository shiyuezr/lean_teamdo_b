package account

import (
	"github.com/bitly/go-simplejson"
	"github.com/kfchen81/beego/vanilla"
	"teamdo/business/task"
)

type User struct {
	vanilla.EntityBase
	Id       int
	Username string
	Password string
	Email    string
	Token    string //身份识别码
	Status   int    //用户状态
	RawData  *simplejson.Json

	TasksUser []*task.Task
}

func init() {
}
