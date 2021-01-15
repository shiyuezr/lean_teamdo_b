package routers

import (
	"teamdo/rest/dev"
	"teamdo/rest/project"

	"github.com/kfchen81/beego/vanilla"
)

func init() {
	vanilla.Router(&dev.BDDReset{})
	vanilla.Router(&project.Project{})
}
