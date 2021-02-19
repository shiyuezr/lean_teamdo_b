package task

import (
	"teamdo/business/account"
	"teamdo/business/lane"
	"teamdo/business/project"
)

type RTask struct {
	Id      int    `json:"id"`
	Content string `json:"content"` //任务内容
	Status  int    `json:"status"`  //任务的完成状态

	User       *account.User    `json:"user"`        //任务执行者
	ParentTask *Task            `json:"parent_task"` //父级任务
	Lane       *lane.Lane       `json:"lane"`        //任务所在泳道
	Project    *project.Project `json:"project"`     //任务所属项目
}
