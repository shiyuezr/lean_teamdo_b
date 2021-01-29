package account

import (
	"context"

	"github.com/bitly/go-simplejson"

	"github.com/kfchen81/beego"
	"github.com/kfchen81/beego/vanilla/middleware"
	"net/http"
)

var gInstance *ContextFactory

type ContextFactory struct {
}

//NewContext 构造含有corp的Context
func (this *ContextFactory) NewContext(ctx context.Context, request *http.Request, userId int, jwtToken string, rawData *simplejson.Json) context.Context {
	platformId := 0
	if rawData != nil {
		platformId, _ = rawData.Get("pid").Int()
	}

	//创建user
	user := new(User)
	user.Model = nil
	user.Id = userId
	user.PlatformId = platformId
	user.RawData = rawData
	user.Ctx = ctx
	ctx = context.WithValue(ctx, "user", user)

	//创建corp
	if rawData != nil {
		jwtType, _ := rawData.Get("type").Int()
		if jwtType == 1 {
			corpId, err := rawData.Get("cid").Int()
			if err == nil {
				corp := new(Corporation)
				corp.Model = nil
				corp.Id = corpId
				corp.Ctx = ctx
				ctx = context.WithValue(ctx, "corp", corp)
			} else {
				beego.Error(err)
			}

		}
	}

	return ctx
}

func init() {
	gInstance = &ContextFactory{}
	beego.Info("433443434", gInstance)
	middleware.SetBusinessContextFactory(gInstance)
}
