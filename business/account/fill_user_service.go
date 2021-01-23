package account

import (
	"context"
	"github.com/kfchen81/beego/vanilla"
)

type FillUserService struct {
	vanilla.ServiceBase
}

func NewFillUserService(ctx context.Context) *FillUserService {
	service := new(FillUserService)
	service.Ctx = ctx
	return service
}

func (this *FillUserService) Fill(users []*User, option vanilla.FillOption) {
	if len(users) == 0 {
		return
	}

	ids := make([]int, 0)
	for _, user := range users {
		ids = append(ids, user.Id)
	}
	return
}
