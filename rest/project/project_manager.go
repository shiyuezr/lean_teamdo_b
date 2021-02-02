package project

import (
	"github.com/kfchen81/beego/vanilla"
	b_project "teamdo/business/project"
)

type ProjectManager struct {
	vanilla.RestResource
}

func (this *ProjectManager) Resource() string {
	return "project.manager"
}

func (this *ProjectManager) GetParameters() map[string][]string {
	return map[string][]string{
		"PUT":    []string{"uid:int", "pid:int"},
		"DELETE": []string{"uid:int", "pid:int"},
	}
}

func (this *ProjectManager) Put() {
	uid, _ := this.GetInt("uid")
	pid, _ := this.GetInt("pid")
	bCtx := this.GetBusinessContext()

	repository := b_project.NewProjectRepository(bCtx)
	project := repository.GetProjectById(pid)
	project.AuthorityVerify()

	respProject := project.AddManager(uid)
	respProject = respProject.AddMember(uid)

	response := vanilla.MakeResponse(vanilla.Map{
		"id": respProject.Id,
	})
	this.ReturnJSON(response)
}

func (this *ProjectManager) Delete() {
	uid, _ := this.GetInt("uid")
	pid, _ := this.GetInt("pid")
	bCtx := this.GetBusinessContext()

	repository := b_project.NewProjectRepository(bCtx)
	project := repository.GetProjectById(pid)
	project.AuthorityVerify()

	respProject := project.DeleteMember(uid)
	respProject = project.DeleteManager(uid)

	response := vanilla.MakeResponse(vanilla.Map{
		"id": respProject.Id,
	})
	this.ReturnJSON(response)
}
