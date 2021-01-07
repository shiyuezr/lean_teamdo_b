package project

import (
	"github.com/kfchen81/beego/vanilla"
	b_project "teamdo/business/project"
)

type Projects struct {
	vanilla.RestResource
}

func (this Projects) Resource() string {
	return "project.projects"
}

func (this *Projects) GetParameters() map[string][]string {
	return map[string][]string{
		"GET": []string{"user_id: int"},
	}
}

func (this *Projects) Get()  {
	bCtx := this.GetBusinessContext()
	userId, _ := this.GetInt("user_id")

	projects := b_project.NewProjectRepository(bCtx).GetProjectsByUserId(userId)
	data := b_project.NewEncodeProjectService(bCtx).EncodeMany(projects)
	response := vanilla.MakeResponse(data)
	this.ReturnJSON(response)
}