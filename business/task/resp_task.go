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

	Operator   *account.RUser    `json:"operator"`    //任务执行者
	ParentTask *RTask            `json:"parent_task"` //父级任务
	Lane       *lane.RLane       `json:"lane"`        //任务所在泳道
	Project    *project.RProject `json:"project"`     //任务所属项目
}
