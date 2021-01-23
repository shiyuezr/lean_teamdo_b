package account

import (
	"context"

	"github.com/kfchen81/beego/vanilla"
)

type EncodeUserService struct {
	vanilla.ServiceBase
}

func NewEncodeUserService(ctx context.Context) *EncodeUserService {
	service := new(EncodeUserService)
	service.Ctx = ctx
	return service
}

//Encode 对单个实体对象进行编码
func (this *EncodeUserService) Encode(user *User) *RUser {
	if user == nil {
		return nil
	}

	return &RUser{
		Id:       user.Id,
		Username: user.Username,
		Password: user.Password,
		Token:    user.Token,
	}
}
