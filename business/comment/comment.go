package comment

import (
	"github.com/kfchen81/beego/vanilla"
	"time"
)

type Comment struct {
	vanilla.EntityBase
	Id              int
	Content         string
	CreatedAt       time.Time
	UsernameComment string
}
