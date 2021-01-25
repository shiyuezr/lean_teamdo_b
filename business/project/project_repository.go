package project

import (
	"context"
	"github.com/kfchen81/beego/vanilla"
)

// todo 暂时当作数据库使用，后期切换mysql
var Project_List = []*Project{
	{
		Id:          1,
		Name:        "pc_development",
		Information: "development",
	},
}

type ProjectRepository struct {
	vanilla.RepositoryBase
}

func (this *ProjectRepository) GetProjectByName(name string) *Project {
	resp_project := &Project{}
	// todo 后期换成从数据库读取数据
	for _, value_project := range Project_List {
		if name == value_project.Name {
			resp_project = value_project
			return resp_project
		}
	}
	return resp_project
}

func (this *ProjectRepository) GetProjectById(id int) *Project {
	resp_project := &Project{}
	// todo 后期换成从数据库读取数据
	for _, value_project := range Project_List {
		if id == value_project.Id {
			resp_project = value_project
			return resp_project
		}
	}
	return resp_project
}

func NewProjectRepository(ctx context.Context) *ProjectRepository {
	repository := new(ProjectRepository)
	repository.Ctx = ctx
	return repository
}
