package project

import "github.com/kfchen81/beego/vanilla"

type ProjectTask struct {
	vanilla.RestResource
}

func (this *ProjectTask) Resource() string {
	return "project.project.task"
}

func (this *ProjectTask) GetParameters() map[string][]string {
	return map[string][]string{
		"GET": []string{"tunnel_id: int"},
		"PUT": []string{
			"tunnel_id: int",
			"title: string",
			"status: bool",
			"remark: string",
			"priority: string",
		},
		"POST": []string{},
	}
}
