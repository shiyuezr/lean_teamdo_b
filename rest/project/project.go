package project

import (
	"github.com/kfchen81/beego/vanilla"
	b_user "teamdo/business/user"
	b_project "teamdo/business/project"
)
type Project struct {
	vanilla.RestResource
}

func (this *Project) Resource() string {
	return "project.project"
}

func (this *Project) GetParameters() map[string][]string {
	return map[string][]string{
		"GET": []string{},
		"PUT": []string{
			"project_name:string",
			"project_detail:string",
		},
		"POST": []string{},
		"DELETE": []string{
			"id:int",
		},
	}
}

func (this *Project)Put()  {
	bCtx:=this.GetBusinessContext()
	userId:= b_user.GetUserFromContext(bCtx).Id
	projectName:=this.GetString("project_name")
	projectDetail:=this.GetString("project_detail")
	project:= b_project.NewProject(bCtx,projectName,projectDetail,userId)
	response := vanilla.MakeResponse(vanilla.Map{
		"id": project.Id,
	})
	this.ReturnJSON(response)
}

func (this *Project)Get()  {
	bCtx:=this.GetBusinessContext()
	userId:= b_user.GetUserFromContext(bCtx).Id
	projects:=b_project.NewProjectRepository(bCtx).GetProjectByUserId(userId)
	if projects==nil {
		panic(vanilla.NewBusinessError("Invalid_projects","获取项目列表为空"))
	}
	encodeProject:=b_project.NewEncodeProjectService(bCtx).EncodeMany(projects)
	response := vanilla.MakeResponse(encodeProject)
	this.ReturnJSON(response)
}