package team

import (
	"github.com/kfchen81/beego"
	"github.com/kfchen81/beego/vanilla"
)

type Team struct {
	vanilla.RestResource
}

func (this *Team) Resource() string {
	return "team.team"
}

func (this *Team) GetParameters() map[string][]string {
	return map[string][]string{
		"GET": []string{
			"id: int",
		},
	}
}

func (this *Team) Get()  {
	beego.Info("初始化的数据的格式入校")
}