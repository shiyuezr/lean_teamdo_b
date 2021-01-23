package comment

import (
	"context"
	"github.com/kfchen81/beego/vanilla"
)

type FillCommentService struct {
	vanilla.ServiceBase
}

func NewFillCommentService(ctx context.Context) *FillCommentService {
	service := new(FillCommentService)
	service.Ctx = ctx
	return service
}
