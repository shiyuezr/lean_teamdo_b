package lane

import (
	"github.com/kfchen81/beego/vanilla"
	"teamdo/business/project"
)

type Lane struct {
	vanilla.EntityBase
	Id   int
	Name string

	Project *project.Project //所属项目
}
