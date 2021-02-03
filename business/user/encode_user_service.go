package user

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

func (this *EncodeUserService) Encode(user *User) *RUser {
	if user == nil {
		return nil
	}

	return &RUser{
		Name: user.UserName,
		Id: user.Id,
	}
}

func (this *EncodeUserService) EncodeMany(users []*User) []*RUser {
	rDatas := make([]*RUser, 0)
	for _, user := range users {
		rDatas = append(rDatas, this.Encode(user))
	}
	return rDatas
}
