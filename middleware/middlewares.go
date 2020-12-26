package middleware

import (
	"github.com/kfchen81/beego"
	"github.com/kfchen81/beego/vanilla/middleware"
)

func init() {
	beego.Notice("register middlewares...")
	beego.InsertFilter("*", beego.BeforeRouter, middleware.ModifyRestMethodFilter)
	beego.InsertFilter("*", beego.BeforeRouter, middleware.JWTAuthFilter)
	beego.InsertFilter("*", beego.BeforeRouter, middleware.RequestHeaderDetectFilter)
}
