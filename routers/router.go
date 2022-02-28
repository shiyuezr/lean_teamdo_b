package routers

import (
	"github.com/kfchen81/beego/vanilla"
	"teamdo/rest/dev"
	"teamdo/rest/lane"
	"teamdo/rest/project"
	"teamdo/rest/project_member"
	"teamdo/rest/user"
)

func init() {
	vanilla.Router(&dev.BDDReset{})
	vanilla.Router(&lane.Lane{})
	vanilla.Router(&project_member.ProjectMember{})
	vanilla.Router(&project.Project{})
	vanilla.Router(&user.User{})
}
