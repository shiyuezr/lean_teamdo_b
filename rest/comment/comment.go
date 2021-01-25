package comment

import (
	"github.com/kfchen81/beego/vanilla"
)

type Comment struct {
	vanilla.RestResource
}

//用return的东西注册路由：vanilla/router.go中注册
func (this *Comment) Resource() string {
	return "comment.comment"
}
