package project_member

import "github.com/kfchen81/beego/orm"

type ProjectMember struct {
	Id int
	ProjectId int
	UserId int
	MemberLevel int
	IsDelete bool
}

func (this *ProjectMember) TableName()  string {
	return "project_member"
}
func init()  {
	orm.RegisterModel(new(ProjectMember))
}