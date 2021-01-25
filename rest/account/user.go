package account

import (
	"github.com/kfchen81/beego/vanilla"
	b_account "teamdo/business/account"
)

type User struct {
	vanilla.RestResource
}

func (this *User) Resource() string {
	return "account_exist.user"
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

	repository := b_account.NewUserRepository(bCtx)
	user := repository.GetUserByInformation(username, password)

	login := b_account.NewLoginService(bCtx)
	user_login := login.Login(user)

	encodeService := b_account.NewEncodeUserService(bCtx)
	respData := encodeService.Encode(user_login)

	response := vanilla.MakeResponse(respData)
	this.ReturnJSON(response)
}
