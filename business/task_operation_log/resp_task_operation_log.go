package task_operation_log

import (
	"teamdo/business/account"
	"teamdo/business/task"
)

type RTaskOperationLog struct {
	Id        int    `json:"id"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`

	Task     *task.RTask    `json:"task"`     //日志所属任务
	Operator *account.RUser `json:"operator"` //日志撰写者
}
