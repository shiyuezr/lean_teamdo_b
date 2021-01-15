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
func (this *EncodeUserService) Encode(user *User) *RUser {
	if user == nil {
		return nil
	}

	return &RUser{
		Id:        user.Id,
		UserName:  user.UserName,
		UserCode:  user.UserCode,
		Age:       user.Age,
		IsDeleted: user.IsDeleted,
		IsEnabled: user.IsEnabled,
		CreateAt:  user.CreateAt,
	}
}
func (this *EncodeUserService) EncodeMany(users []*User) []*RUser {
	rDatas := make([]*RUser, 0)

	for _, user := range users {
		rDatas = append(rDatas, this.Encode(user))
	}

	return rDatas
}
