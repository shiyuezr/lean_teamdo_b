package log

import (
	"github.com/kfchen81/beego/orm"
	"time"
)

type Log struct {
	Id        int
	Content   string
	TaskId    int       //foreign key Task
	UserId    int       //foreign key User
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
}

func (self *Log) TableName() string {
	return "log_log"
}

func init() {
	orm.RegisterModel(new(Log))
}
