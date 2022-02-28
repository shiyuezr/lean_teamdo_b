package project

import (
	"context"
	"github.com/kfchen81/beego"
	"github.com/kfchen81/beego/vanilla"
	models "teamdo/models/project"
	members "teamdo/models/project_member"
)

type Project struct {
	vanilla.EntityBase
	Id int
	Name string
	ManagerId int
	Detail string
	IsDelete bool
}

func NewProject(ctx context.Context, name string,detail string,managerId int) *Project {
	o:=vanilla.GetOrmFromContext(ctx)
	o.Begin()
	model:=models.Project{}
	model.Name=name
	model.Detail=detail
	model.IsDelete=false
	model.ManagerId=managerId//Format(common.TIME_LAYOUT)
	id,err:=o.Insert(&model)
	if err!=nil {
		o.Rollback()
		beego.Error(err)
		panic(vanilla.NewBusinessError("create_project_failed","创建项目失败"))
	}
	model.Id=int(id)
	member:=members.ProjectMember{}
	member.UserId=managerId
	member.ProjectId=model.Id
	member.MemberLevel=999
	_,memberErr:=o.Insert(&member)
	if memberErr!=nil {
		o.Rollback()
		beego.Error(memberErr)
		panic(vanilla.NewBusinessError("create_project_add_member_failed","创建项目中添加成员失败"))
	}
	o.Commit()
	return NewProjectFromDbModel(ctx,&model)
}


func  NewProjectFromDbModel(ctx context.Context, dbModel *models.Project) *Project {
	instance := new(Project)
	instance.Ctx = ctx
	instance.Id = dbModel.Id
	instance.Name=dbModel.Name
	instance.Detail=dbModel.Detail
	instance.ManagerId = dbModel.ManagerId
	instance.IsDelete = dbModel.IsDelete
	return instance
}