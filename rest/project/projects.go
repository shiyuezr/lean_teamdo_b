package project

import (
	"github.com/kfchen81/beego/vanilla"
)

type Projects struct {
	vanilla.RestResource
}

func (this Projects) Resource() string {
	return "project.projects"
}

func (this *Projects) GetParameters() map[string][]string {
	return map[string][]string{
		"GET": []string{"user_id: int"},
	}
}
