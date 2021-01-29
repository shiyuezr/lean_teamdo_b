package project

import (
	"github.com/kfchen81/beego/vanilla"
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
		"GET": []string{"id:int", "?with_options:json"},
		"PUT": []string{
			"name:string",
			"content:string",
		},
		"POST": []string{
			"id:int",
			"name:string",
			"content:string",
			"status:int",
		},
		"DELETE": []string{"id:int"},
	}
}

func (this *Project) Get() {
	id, _ := this.GetInt("id")
	bCtx := this.GetBusinessContext()

	repository := b_project.NewProjectRepository(bCtx)
	project := repository.GetProjectById(id)
	if project == nil {
		panic(vanilla.NewBusinessError("project_not_exist", "项目不存在"))
	}

	b_project.NewFillProjectService(bCtx).FillOne(project, this.GetFillOptions("with_options"))

	encodeService := b_project.NewEncodeProjectService(bCtx)
	respData := encodeService.Encode(project)
	response := vanilla.MakeResponse(respData)
	this.ReturnJSON(response)
}

func (this *Project) Put() {
	name := this.GetString("name")
	content := this.GetString("content")
	bCtx := this.GetBusinessContext()

	project := b_project.NewProject(bCtx, name, content)
	response := vanilla.MakeResponse(vanilla.Map{
		"id": project.Id,
	})
	this.ReturnJSON(response)
}

func (this *Project) Delete() {
	id, _ := this.GetInt("id")
	bCtx := this.GetBusinessContext()

	project := b_project.NewProjectRepository(bCtx).GetProjectById(id)
	project.Delete()

	response := vanilla.MakeResponse(vanilla.Map{
		"id": id,
	})
	this.ReturnJSON(response)
}

func (this *Project) Post() {
	id, _ := this.GetInt("id")
	name := this.GetString("name")
	content := this.GetString("content")
	status, _ := this.GetInt("status")
	bCtx := this.GetBusinessContext()
	jwt := this.Ctx.Request.Header.Get("Authorization")

	repository := b_project.NewProjectRepository(bCtx)
	project := repository.GetProjectById(id)
	project.AuthorityVerify(jwt)
	_ = project.Update(name, content, status)

	response := vanilla.MakeResponse(vanilla.Map{})
	this.ReturnJSON(response)
}
