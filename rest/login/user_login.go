package login

import (
	"encoding/base64"
	"fmt"
	"github.com/kfchen81/beego/vanilla"
	b_account "teamdo/business/account"
)

type LoginUser struct {
	vanilla.RestResource
}

func (this *LoginUser) Resource() string {
	return "login.logined.user"
}

func (this *LoginUser) GetParameters() map[string][]string {
	return map[string][]string{
		"PUT": []string{
			"username:string",
			"password:string",
		},
	}
}

// 登录
func (this *LoginUser) Put() {
	username := this.GetString("username")
	password := this.GetString("password")
	filters := vanilla.Map{
		"username": username,
		"password": base64.StdEncoding.EncodeToString([]byte(password)),
	}
	bCtx := this.GetBusinessContext()

	repository := b_account.NewUserRepository(bCtx)
	loginUsers := repository.GetByFilters(filters)
	if len(loginUsers) == 0 {
		panic(vanilla.NewSystemError("user:login_fail", fmt.Sprintf("登录失败，用户名或密码错误")))
	}
	respUser := loginUsers[0].Login()

	encodeService := b_account.NewEncodeUserService(bCtx)
	respData := encodeService.Encode(respUser)

	response := vanilla.MakeResponse(respData)
	this.ReturnJSON(response)
}
