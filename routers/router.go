package routers

import (
	"github.com/kfchen81/beego/vanilla"
	"teamdo/rest/account"
	"teamdo/rest/comment"
	"teamdo/rest/dev"
	"teamdo/rest/project"
)

func init() {
	vanilla.Router(&dev.BDDReset{})
	vanilla.Router(&comment.Comment{})
	vanilla.Router(&account.User{})
	vanilla.Router(&project.Project{})
}
