package project

import "github.com/kfchen81/beego/orm"

type Project struct {
	Id      int
	Name    string
	Content string
	Status  int `orm:"default(1)"` //完成状态
}

func (self *Project) TableName() string {
	return "project_project"
}

// 项目和管理员多对多表
type ProjectToAdministrators struct {
	Id              int
	ProjectId       int
	AdministratorId int
}

// 项目和参与者多对多表
type ProjectToParticipants struct {
	Id            int
	ProjectId     int
	ParticipantId int
}

func init() {
	orm.RegisterModel(new(Project))
	orm.RegisterModel(new(ProjectToAdministrators))
	orm.RegisterModel(new(ProjectToParticipants))
}
