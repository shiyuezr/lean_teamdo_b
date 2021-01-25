package account

import (
	"context"
	"github.com/kfchen81/beego/vanilla"
)

type LoginService struct {
	vanilla.ServiceBase
}

// Login 给予jwt
func (this *LoginService) Login(user *User) *User {
	if user.Id != 0 {
		jwttoken := vanilla.EncodeJWT(vanilla.Map{
			"type": 2,
			"uid":  user.Id,
		})
		user.Token = jwttoken
		return user
	}
	return user
}

func NewLoginService(ctx context.Context) *LoginService {
	service := new(LoginService)
	service.Ctx = ctx
	return service
}
