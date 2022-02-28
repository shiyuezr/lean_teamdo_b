package project_member

import (
	"github.com/kfchen81/beego/vanilla"
	b_project_member "teamdo/business/project_member"
)

type ProjectMember struct {
	vanilla.RestResource
}

func (this *ProjectMember) Resource() string {
	return "project_member.project_member"
}

func (this *ProjectMember) GetParameters() map[string][]string {
	return map[string][]string{
		"GET": []string{},
		"PUT": []string{
			"project_id:int",
			"user_id:int",
		},
		"POST": []string{},
		"DELETE": []string{
			"id:int",
		},
	}
}

func (this *ProjectMember)Put()  {
	bCtx:=this.GetBusinessContext()
	projectId,_:=this.GetInt("project_id")
	userId,_:=this.GetInt("user_id")
	projectMemberId:= b_project_member.NewProjectMember(bCtx,projectId,userId)
	response := vanilla.MakeResponse(vanilla.Map{
		"id":projectMemberId.Id,
	})
	this.ReturnJSON(response)
}