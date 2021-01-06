package project

import (
	"github.com/kfchen81/beego/vanilla"
)

type ProjectTunnel struct {
	vanilla.RestResource
}

func (this *ProjectTunnel) Resource() string {
	return "project.project_tunnel"
}

func (this *ProjectTunnel) GetParameters() map[string][]string {
	return map[string][]string{
		"GET": []string{"project_id: int"},
		"PUT": []string{"project_id: int", "title: string", "manager_id: int"},
		"POST": []string{"id: int"},
		"DELETE": []string{"id: int"},
	}
}

func (this *ProjectTunnel) Get()  {

}

func (this *ProjectTunnel) Put()  {

}

func (this *ProjectTunnel) Post()  {

}

func (this *ProjectTunnel) Delete()  {

}
