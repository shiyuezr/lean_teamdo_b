package comment

import (
	"github.com/kfchen81/beego/vanilla"
	"teamdo/business/account"
	"teamdo/business/task"
	"time"
)

type Comment struct {
	vanilla.EntityBase
	Id        int
	Content   string
	CreatedAt time.Time

	User *account.User //评论人
	Task *task.Task    //所属任务
}
