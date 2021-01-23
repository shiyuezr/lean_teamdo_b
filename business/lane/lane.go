package lane

import (
	"github.com/kfchen81/beego/vanilla"
	"teamdo/business/task"
)

type Lane struct {
	vanilla.EntityBase
	Id   int
	Name string

	TasksLane []*task.Task
}
