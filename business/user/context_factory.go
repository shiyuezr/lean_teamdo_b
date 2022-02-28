package user

import (
	"context"

	"github.com/bitly/go-simplejson"

	"github.com/kfchen81/beego/vanilla/middleware"
	"net/http"
)

var gInstance *ContextFactory

type ContextFactory struct {
}

//NewContext 构造含有corp的Context
func (this *ContextFactory) NewContext(ctx context.Context, request *http.Request, userId int, jwtToken string, rawData *simplejson.Json) context.Context {
	//platformId := 0
	if rawData != nil {
		//platformId, _ = rawData.Get("pid").Int()

		sourceService, err := rawData.Get("__source_service").String()
		if err != nil {
			sourceService = "tippler"
		}
		ctx = context.WithValue(ctx, "source_service", sourceService)
	}

	//创建user
	user := new(User)
	user.Model = nil
	user.Id = userId
	//user.PlatformId = platformId
	//user.RawData = rawData
	user.Ctx = ctx
	ctx = context.WithValue(ctx, "user", user)

	////创建corp
	//if rawData != nil {
	//	jwtType, _ := rawData.Get("type").Int()
	//	if jwtType == 1 {
	//		corpId, err := rawData.Get("cid").Int()
	//		if err == nil {
	//			ctx = context.WithValue(ctx, "cid", corpId)
	//			// corp user
	//			corpUserId, _ := rawData.Get("uid").Int()
	//			ctx = context.WithValue(ctx, "uid", corpUserId)
	//
	//			// corp related user
	//			relatedUserId, _ := rawData.Get("cru_id").Int()
	//			ctx = context.WithValue(ctx, "cru_id", relatedUserId)
	//		} else {
	//			beego.Error(err)
	//		}
	//	} else if jwtType == 3{
	//		corpUserData := rawData.Get("corp_user")
	//		corpId, _ := corpUserData.Get("cid").Int()
	//		ctx = context.WithValue(ctx, "cid", corpId)
	//	} else {
	//		corpId, err := rawData.Get("cid").Int()
	//		if err == nil {
	//			ctx = context.WithValue(ctx, "cid", corpId)
	//		}
	//	}
	//}

	return ctx
}

func init() {
	gInstance = &ContextFactory{}
	middleware.SetBusinessContextFactory(gInstance)
}
