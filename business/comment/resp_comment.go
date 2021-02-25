package comment

import (
	"teamdo/business/account"
	"teamdo/business/task"
)

type RComment struct {
	Id        int            `json:"id"`
	Content   string         `json:"content"`
	CreatedAt string         `json:"created_at"`
	TaskId    int            `json:"task_id"`
	UserId    int            `json:"user_id"`
	User      *account.RUser `json:"user"` //评论人
	Task      *task.RTask    `json:"task"` //所属任务
}
