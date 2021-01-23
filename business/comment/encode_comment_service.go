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

////Encode 对单个实体对象进行编码
//func (this *EncodeCommentService) Encode(Comment *Comment) *RComment {
//	if tag == nil {
//		return nil
//	}
//
//	return &RTag{
//		Id:        tag.Id,
//		Name:      tag.Name,
//		IsEnabled: tag.IsEnabled,
//		IsDeleted: tag.IsDeleted,
//		CreatedAt: tag.CreatedAt.Format("2006-01-02 15:04:05"),
//	}
//}
