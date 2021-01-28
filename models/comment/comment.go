package comment

import (
	"github.com/kfchen81/beego/orm"
	"time"
)

type Comment struct {
	Id        int
	Content   string
	UserId    int       //foreign key User
	TaskId    int       //foreign key Task
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
}

func (self *Comment) TableName() string {
	return "comment_comment"
}

func init() {
	orm.RegisterModel(new(Comment))
}
