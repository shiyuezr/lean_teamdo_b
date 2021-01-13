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
		"GET": []string{
			"id:int",
		},
		"PUT": []string{
			"user_id:int",
			"project_name:string",
		},
	}
}

func (this *Project) Get()  {
	bCtx := this.GetBusinessContext()

	id, _ := this.GetInt("id")
	project := b_project.NewProjectRepository(bCtx).GetProjectById(id)
	if project == nil {
		panic(vanilla.NewBusinessError("invalid_projectId", "项目不存在"))
	}
	rProject := b_project.NewEncodeProjectService(bCtx).Encode(project)

	response := vanilla.MakeResponse(vanilla.Map{
		"project": rProject,
	})
	this.ReturnJSON(response)
}


func (this *Project) Put()  {
	bCtx := this.GetBusinessContext()

	userId, _ := this.GetInt("user_id")
	projectName := this.GetString("project_name")
	b_project.NewProject(bCtx, userId, projectName)

	response := vanilla.MakeResponse(vanilla.Map{})
	this.ReturnJSON(response)
}