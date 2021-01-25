package project

import (
	"context"
	"github.com/kfchen81/beego/vanilla"
)

type ProjectOperationService struct {
	vanilla.RepositoryBase
}

func (this *ProjectOperationService) ProjectInsert(project *Project) int {
	Project_List = append(Project_List, project)
	return project.Id
}

// ProjectDelete 项目的删除操作，根据id删除项目，返回0代表删除失败
func (this *ProjectOperationService) ProjectDelete(id int) int {
	count := -1
	for index, project_value := range Project_List {
		if project_value.Id == id {
			count = index
		}
	}
	Project_List = append(Project_List[:count], Project_List[count+1:]...)
	return count + 1
}

func (this *ProjectOperationService) ProjectModify(id int, name string, information string) int {
	for _, project_value := range Project_List {
		if project_value.Id == id {
			project_value.Name = name
			project_value.Information = information
		}
	}
	return id
}

func NewProjectOperationService(ctx context.Context) *ProjectOperationService {
	repository := new(ProjectOperationService)
	repository.Ctx = ctx
	return repository
}
