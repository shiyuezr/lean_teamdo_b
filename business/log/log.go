package log

import (
	"github.com/kfchen81/beego/vanilla"
	"teamdo/business/account"
	"teamdo/business/task"
	"time"
)

type Log struct {
	vanilla.EntityBase
	Id        int
	CreatedAt time.Time
	Content   string

	Task *task.Task    //日志所属任务
	User *account.User //日志撰写者
}
