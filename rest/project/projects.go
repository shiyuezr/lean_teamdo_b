package project

import (
	"github.com/kfchen81/beego/vanilla"
	b_project "teamdo/business/project"
)

type Projects struct {
	vanilla.RestResource
}

// Resource 项目列表
func (this *Projects) Resource() string {
	return "project.projects"
}

func (this *Projects) GetParameters() map[string][]string {
	return map[string][]string{
		"GET": []string{"?filters:json"},
	}
}

func (this *Projects) Get() {
	bCtx := this.GetBusinessContext()
	filters := vanilla.ConvertToBeegoOrmFilter(this.GetFilters())

	projects := b_project.NewProjectRepository(bCtx).GetByFilters(filters)

	response := vanilla.MakeResponse(b_project.NewEncodeProjectService(bCtx).EncodeMany(projects))
	this.ReturnJSON(response)
}
