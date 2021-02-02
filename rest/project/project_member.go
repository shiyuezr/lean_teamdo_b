package project

import (
	"github.com/kfchen81/beego/vanilla"
	b_project "teamdo/business/project"
)

type ProjectMember struct {
	vanilla.RestResource
}

func (this *ProjectMember) Resource() string {
	return "project.member"
}

func (this *ProjectMember) GetParameters() map[string][]string {
	return map[string][]string{
		"PUT":    []string{"uid:int", "pid:int"},
		"DELETE": []string{"uid:int", "pid:int"},
	}
}

func (this *ProjectMember) Put() {
	uid, _ := this.GetInt("uid")
	pid, _ := this.GetInt("pid")
	bCtx := this.GetBusinessContext()

	project := b_project.NewProjectRepository(bCtx).GetProjectById(pid)
	project.AuthorityVerify()

	respProject := project.AddMember(uid)

	response := vanilla.MakeResponse(vanilla.Map{
		"id": respProject.Id,
	})
	this.ReturnJSON(response)
}

func (this *ProjectMember) Delete() {
	uid, _ := this.GetInt("uid")
	pid, _ := this.GetInt("pid")
	bCtx := this.GetBusinessContext()

	repository := b_project.NewProjectRepository(bCtx)
	project := repository.GetProjectById(pid)
	project.AuthorityVerify()

	respProject := project.DeleteMember(uid)

	response := vanilla.MakeResponse(vanilla.Map{
		"id": respProject.Id,
	})
	this.ReturnJSON(response)
}
