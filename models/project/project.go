package project

import "github.com/kfchen81/beego/orm"

type Project struct {
	Id int
	Name string
	ManagerId int
	Detail string
	IsDelete bool
}


func (this *Project) TableName()  string {
	return "project_project"
}
func init()  {
	orm.RegisterModel(new(Project))
}