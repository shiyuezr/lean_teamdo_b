package project

import (
	"github.com/kfchen81/beego/vanilla"
)

type Project struct {
	vanilla.RestResource
}

//用return的东西注册路由：vanilla/router.go中注册
func (this *Project) Resource() string {
	return "project.project"
}

func (this *Project) get() {

}
