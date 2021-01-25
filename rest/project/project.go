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
		"GET": []string{"id:int"},
		"PUT": []string{
			"name:string",
			"information:string",
		},
		"POST": []string{
			"name:string",
			"information:string",
		},
		"DELETE": []string{"id:int"},
	}
}

func (this *Project) Get() {
	id,_ := this.GetInt("id")
	bCtx := this.GetBusinessContext()

	repository := b_project.NewProjectRepository(bCtx)
	project := repository.GetProjectById(id)

	encodeService := b_project.NewEncodeProjectService(bCtx)
	respData := encodeService.Encode(project)

	response := vanilla.MakeResponse(respData)
	this.ReturnJSON(response)
}

func (this *Project) Put() {
	name := this.GetString("name")
	information := this.GetString("information")

	bCtx := this.GetBusinessContext()

	project := b_project.NewProject(bCtx, name, information)

	operation := b_project.NewProjectOperationService(bCtx)
	resp_id := operation.ProjectInsert(project)

	response := vanilla.MakeResponse(vanilla.Map{
		"id": resp_id,
	})
	this.ReturnJSON(response)
}

func (this *Project) Delete() {
	id, _ := this.GetInt("id")

	bCtx := this.GetBusinessContext()

	operation := b_project.NewProjectOperationService(bCtx)
	resp_id := operation.ProjectDelete(id)

	response := vanilla.MakeResponse(vanilla.Map{
		"id": resp_id,
	})
	this.ReturnJSON(response)
}

func (this *Project) Post() {
	name := this.GetString("name")
	information := this.GetString("information")

	bCtx := this.GetBusinessContext()

	repository := b_project.NewProjectRepository(bCtx)
	project := repository.GetProjectByName(name)

	operation := b_project.NewProjectOperationService(bCtx)
	resp_id := operation.ProjectModify(project.Id, name, information)

	response := vanilla.MakeResponse(vanilla.Map{
		"id": resp_id,
	})
	this.ReturnJSON(response)
}
