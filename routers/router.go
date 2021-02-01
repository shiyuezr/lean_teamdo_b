package routers

import (
	"github.com/kfchen81/beego/vanilla"
	"teamdo/rest/comment"
	"teamdo/rest/dev"
	"teamdo/rest/login"
	"teamdo/rest/project"
)

func init() {
	vanilla.Router(&dev.BDDReset{})
	vanilla.Router(&comment.Comment{})
	vanilla.Router(&login.LoginUser{})
	vanilla.Router(&project.Project{})
	vanilla.Router(&project.Projects{})
	vanilla.Router(&project.ProjectAddManager{})
}
