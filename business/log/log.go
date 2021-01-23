package log

import (
	"github.com/kfchen81/beego/vanilla"
	"time"
)

type Log struct {
	vanilla.EntityBase
	Id          int
	CreatedAt   time.Time
	Content     string
	UsernameLog string
}
