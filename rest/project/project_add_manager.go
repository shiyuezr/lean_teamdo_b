package project

import (
	"github.com/kfchen81/beego/vanilla"
	b_project "teamdo/business/project"
)

type ProjectAddManager struct {
	vanilla.RestResource
}

func (this *ProjectAddManager) Resource() string {
	return "project.manager"
}

func (this *ProjectAddManager) GetParameters() map[string][]string {
	return map[string][]string{
		"PUT":    []string{"uid:int", "pid:int"},
		"DELETE": []string{"id:int"},
	}
}

func (this *ProjectAddManager) Put() {
	uid, _ := this.GetInt("uid")
	pid, _ := this.GetInt("pid")
	bCtx := this.GetBusinessContext()

	repository := b_project.NewProjectRepository(bCtx)
	project := repository.GetProjectById(pid)
	project.AuthorityVerify()

	respProject := project.AddManager(uid)
	// todo 添加管理员同时也要添加参与者，因为参与者也是管理员
	//respProject = respProject.Addmember(uid)

	response := vanilla.MakeResponse(vanilla.Map{
		"id": respProject.Id,
	})
	this.ReturnJSON(response)
}
