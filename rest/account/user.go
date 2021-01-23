package account

import (
	"github.com/kfchen81/beego/vanilla"
	"teamdo/business/account"
)

type User struct {
	vanilla.RestResource
}

//用return的东西注册路由：vanilla/router.go中注册
func (this *User) Resource() string {
	return "account.user"
}

//限制传入参数
func (this *User) GetParameters() map[string][]string {
	return map[string][]string{
		"GET": []string{"id:int"},
		"PUT": []string{
			"id:int",
			"username:string",
			"password:string",
		},
		"POST": []string{
			"id:string",
			//"id:int",
			"username:string",
			"password:string",
		},
		"DELETE": []string{"id:int"},
	}
}

func (this *User) Get() {

}

//登录
func (this *User) Post() {
	id, _ := this.GetInt("id")
	username := this.GetString("username")
	password := this.GetString("password")

	bCtx := this.GetBusinessContext() //这个只用来初始化服务么

	verify := account.NewVerifyUserService(bCtx)
	jwt := verify.Verify(username, password, id)

	//使用bCtx构建UserRepository对象
	repository := account.NewUserRepository(bCtx)
	user := repository.GetUser(id)
	user.Token = jwt

	//fillService := account.NewFillUserService(bCtx)
	//fillService.Fill([]*account.User{user}, vanilla.FillOption{})

	encodeService := account.NewEncodeUserService(bCtx)
	respData := encodeService.Encode(user)

	response := vanilla.MakeResponse(respData)
	this.ReturnJSON(response)

	//Comment := repository.GetComment(id)
}

//func (this *User) post() {
//	username := this.GetString("username")
//	password := this.GetString("password")
//	bCtx := this.GetBusinessContext()
//}
