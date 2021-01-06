package project

import (
	"github.com/kfchen81/beego/vanilla"
	"time"
)

type Tunnel struct {
	vanilla.EntityBase
	Id			int
	Title 		string
	ProjectId   int
	IsDelete    bool
	CreateAt    time.Time

	Task        []*Task
}
