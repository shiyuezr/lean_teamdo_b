package account

import (
	"context"
	"github.com/kfchen81/beego/vanilla"
)

type EncodeUserService struct {
	vanilla.ServiceBase
}

// Encode 对单个实体对象进行编码
func (this *EncodeUserService) Encode(user *User) *RUser {
	if user == nil {
		return nil
	}
	return &RUser{
		Id:       user.Id,
		Username: user.Username,
		Token:    user.Token,
	}
}

// EncodeMany 对实体对象进行批量编码
func (this *EncodeUserService) EncodeMany(users []*User) []*RUser {
	rDatas := make([]*RUser, 0)
	for _, user := range users {
		rDatas = append(rDatas, this.Encode(user))
	}

	return rDatas
}

func NewEncodeUserService(ctx context.Context) *EncodeUserService {
	service := new(EncodeUserService)
	service.Ctx = ctx
	return service
}
