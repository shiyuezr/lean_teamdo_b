package project_member

import (
	"context"
	"github.com/kfchen81/beego/vanilla"
	"github.com/kfchen81/beego"
	models "teamdo/models/project_member"
)
type ProjectMember struct {
	vanilla.EntityBase
	Id int
	ProjectId int
	UserId int
	IsDelete bool
}

func NewProjectMember(ctx context.Context, projectId int,userId int) *ProjectMember {
	o:=vanilla.GetOrmFromContext(ctx)
	model:=models.ProjectMember{}
	model.ProjectId=projectId
	model.UserId=userId
	model.MemberLevel=1
	model.IsDelete=false
	id,err:=o.Insert(&model)
	if err!=nil {
		beego.Error(err)
		panic(vanilla.NewBusinessError("create_project_failed","创建项目失败"))
	}
	model.Id=int(id)
	return NewProjectMemberFromDbModel(ctx,&model)
}

func  NewProjectMemberFromDbModel(ctx context.Context, dbModel *models.ProjectMember) *ProjectMember {
	instance := new(ProjectMember)
	instance.Ctx = ctx
	instance.Id = dbModel.Id
	instance.ProjectId=dbModel.ProjectId
	instance.UserId=dbModel.UserId
	instance.IsDelete = dbModel.IsDelete
	return instance
}

//func  NewProjectMembersFromDbModels(ctx context.Context, dbModels []*models.ProjectMember) []*ProjectMember {
//
//	var ProjectMembers []*ProjectMember
//	for k,_:= range dbModels{
//		instance := new(ProjectMember)
//		instance.Ctx = ctx
//		instance.Id = dbModels[k].Id
//		instance.ProjectId=dbModels[k].ProjectId
//		instance.UserId=dbModels[k].UserId
//		instance.IsDelete = dbModels[k].IsDelete
//		ProjectMembers = append(ProjectMembers,instance)
//	}
//	return ProjectMembers
//}