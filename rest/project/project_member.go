package project

import (
	"github.com/kfchen81/beego/vanilla"
	b_project "teamdo/business/project"
)

type Member struct {
	vanilla.RestResource
}

func (this *Member) Resource() string {
	return "project.member"
}

func (this *Member) GetParameter() map[string][]string {
	return map[string][]string{
		"PUT": []string{"project_id:int", "user_id:int"},
		"DELETE": []string{"project_id:int", "user_id:int"},
	}
}

func (this *Member) Put()  {
	bCtx := this.GetBusinessContext()

	projectId, _ := this.GetInt("project_id")
	userId, _ := this.GetInt("user_id")
	project := b_project.NewProjectRepository(bCtx).GetProjectById(projectId)
	project.AddMember(userId)

	response := vanilla.MakeResponse(vanilla.Map{})
	this.ReturnJSON(response)
}

func (this *Member) Delete()  {
	bCtx := this.GetBusinessContext()

	projectId, _ := this.GetInt("project_id")
	userId, _ := this.GetInt("user_id")
	project := b_project.NewProjectRepository(bCtx).GetProjectById(projectId)
	project.DeleteMember(userId)

	response := vanilla.MakeResponse(vanilla.Map{})
	this.ReturnJSON(response)
}
