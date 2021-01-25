package comment

import (
	"context"
	"github.com/kfchen81/beego/vanilla"
)

type EncodeCommentService struct {
	vanilla.ServiceBase
}

func NewEncodeCommentService(ctx context.Context) *EncodeCommentService {
	service := new(EncodeCommentService)
	service.Ctx = ctx
	return service
}
