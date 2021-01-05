package project

import (
	"context"
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

func NewProjectTunnelForModel(ctx context.Context)  {

}