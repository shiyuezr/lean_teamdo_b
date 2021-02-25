package routers

import (
	"github.com/kfchen81/beego/vanilla"
	"teamdo/rest/comment"
	"teamdo/rest/dev"
	"teamdo/rest/lane"
	"teamdo/rest/login"
	"teamdo/rest/project"
	"teamdo/rest/task"
)

func init() {
	vanilla.Router(&dev.BDDReset{})
	vanilla.Router(&comment.Comment{})
	vanilla.Router(&comment.Comments{})
	vanilla.Router(&login.LoginUser{})
	vanilla.Router(&project.Project{})
	vanilla.Router(&project.Projects{})
	vanilla.Router(&project.ProjectManager{})
	vanilla.Router(&project.ProjectMember{})
	vanilla.Router(&lane.Lane{})
	vanilla.Router(&lane.Lanes{})
	vanilla.Router(&task.Task{})
	vanilla.Router(&task.Tasks{})
	vanilla.Router(&task.TaskOperationLog{})
	vanilla.Router(&task.TaskOperationLogs{})
}
