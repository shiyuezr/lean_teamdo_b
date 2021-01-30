package account

import (
	"github.com/kfchen81/beego/vanilla"
	b_account "teamdo/business/account"
)

type User struct {
	vanilla.RestResource
}

func (this *User) Resource() string {
	return "account.logined.user"
}

func (this *User) GetParameters() map[string][]string {
	return map[string][]string{
		"PUT": []string{
			"username:string",
			"password:string",
		},
	}
}

// 登录
func (this *User) Put() {
	username := this.GetString("username")
	password := this.GetString("password")
	bCtx := this.GetBusinessContext()

	loginUser := b_account.NewUserFromLoginInfo(bCtx, username, password)
	respUser := loginUser.Login()

	encodeService := b_account.NewEncodeUserService(bCtx)
	respData := encodeService.Encode(respUser)

	response := vanilla.MakeResponse(respData)
	this.ReturnJSON(response)
}
