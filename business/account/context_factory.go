package account

import (
	"context"
	"github.com/bitly/go-simplejson"
	"github.com/kfchen81/beego/vanilla/middleware"
	"net/http"
)

var gInstance *ContextFactory

type ContextFactory struct {
}

// NewContext
func (this *ContextFactory) NewContext(ctx context.Context, request *http.Request, userId int, jwtToken string, rawData *simplejson.Json) context.Context {
	user := new(User)
	user.Model = nil
	user.Id = userId
	user.RawData = rawData
	user.Ctx = ctx
	ctx = context.WithValue(ctx, "user", user)
	return ctx
}

func init() {
	gInstance = &ContextFactory{}
	middleware.SetBusinessContextFactory(gInstance)
}
