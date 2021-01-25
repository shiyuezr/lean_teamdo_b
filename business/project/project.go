package project

import (
	"context"
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

// todo 后期换成数据库查询，ctx先预设
func NewProject(ctx context.Context, name string, information string) *Project {
	project := &Project{}
	// todo 这里的id，应该是数据库生成的自增id，这里先用2替代，后期换
	project.Id = 2
	project.Name = name
	project.Information = information
	return project
}
