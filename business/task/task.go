package task

import (
	"github.com/kfchen81/beego/vanilla"
	"teamdo/business/account"
	"teamdo/business/lane"
	"teamdo/business/project"
)

type Task struct {
	vanilla.EntityBase
	Id      int
	Content string //任务内容
	Status  int    //任务的完成状态

	User       *account.User    //任务执行者
	ParentTask *Task            //父级任务
	lane       *lane.Lane       //任务所在泳道
	Project    *project.Project //任务所属项目
}
