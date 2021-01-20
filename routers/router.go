package routers

import (
	"github.com/kfchen81/beego/vanilla"
	"teamdo/rest/tunnel"
	"teamdo/rest/dev"
	"teamdo/rest/project"
	"teamdo/rest/task"
	"teamdo/rest/user"
)

func init() {
	vanilla.Router(&dev.BDDReset{})

	vanilla.Router(&project.Project{})
	vanilla.Router(&project.Projects{})
	vanilla.Router(&project.Member{})
	vanilla.Router(&project.Members{})

	vanilla.Router(&task.Task{})
	vanilla.Router(&task.TaskExecutor{})
	vanilla.Router(&task.TaskTitle{})
	vanilla.Router(&task.TaskPriority{})
	vanilla.Router(&task.FinishTask{})

	vanilla.Router(&tunnel.Tunnel{})
	vanilla.Router(&tunnel.Tunnels{})

	vanilla.Router(&user.User{})
}
