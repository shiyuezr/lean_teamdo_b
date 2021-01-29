package account

import (
	"context"
	"github.com/kfchen81/beego/vanilla"
)

type Corporation struct {
	vanilla.EntityBase
	Id int
}

func NewCorpFromOnlyId(ctx context.Context, id int) *Corporation {
	instance := new(Corporation)
	instance.Ctx = ctx
	instance.Model = nil
	instance.Id = id
	return instance
}

//NewContext 构造含有corp的Context
func NewCorpContext(ctx context.Context, corpId int) context.Context {
	corp := new(Corporation)
	corp.Model = nil
	corp.Id = corpId

	corp.Ctx = ctx

	ctx = context.WithValue(ctx, "corp", corp)
	return ctx
}

func GetCorpFromContext(ctx context.Context) *Corporation {
	user := ctx.Value("corp").(*Corporation)
	return user
}

func (this *Corporation) GetId() int {
	return this.Id
}

// IsEnableGreenContentCheck: 是否启用内容审查
func (this *Corporation) IsEnableGreenContentCheck() bool {
	return false
}

func init() {
}
