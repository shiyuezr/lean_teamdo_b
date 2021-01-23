package project

import (
	"github.com/kfchen81/beego/vanilla"
	"teamdo/business/account"
	"teamdo/business/lane"
	"teamdo/business/task"
)

type Project struct {
	vanilla.EntityBase
	Id          int
	Name        string
	Information string
	Status      int //完成状态

	Administrators []*account.User //管理员
	Participant    []*account.User //参与者
	Task           []*task.Task
	Lane           []*lane.Lane
}
