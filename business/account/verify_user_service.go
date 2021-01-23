package account

import (
	"context"
	"github.com/kfchen81/beego/vanilla"
)

type VerifyUserService struct {
	vanilla.ServiceBase
}

func NewVerifyUserService(ctx context.Context) *VerifyUserService {
	service := new(VerifyUserService)
	service.Ctx = ctx
	return service
}

func (this *VerifyUserService) Verify(username string, password string, id int) string {
	//filters := vanilla.Map{
	//	"username": username,
	//	"password": password,
	//}
	//o := vanilla.GetOrmFromContext(this.Ctx)
	//qs :=o.QueryTable(&account.User{})//model表

	if username == "tom" && password == "123456" {
		var jwtoken string = ""
		jwtoken = vanilla.EncodeJWT(map[string]interface{}{
			"type": 1,
			"uid":  id,
		})
		return jwtoken
	}
	return "error"

}
