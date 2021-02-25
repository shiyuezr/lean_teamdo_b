package comment

import (
	"context"
	"github.com/kfchen81/beego/vanilla"
	b_account "teamdo/business/account"
	b_task "teamdo/business/task"
)

type EncodeCommentService struct {
	vanilla.ServiceBase
}

// Encode 对单个实体对象进行编码
func (this *EncodeCommentService) Encode(comment *Comment) *RComment {
	if comment == nil {
		return nil
	}
	rUser := b_account.NewEncodeUserService(this.Ctx).Encode(comment.User)
	rTask := b_task.NewEncodeTaskService(this.Ctx).Encode(comment.Task)

	encodedComment := &RComment{
		Id:        comment.Id,
		Content:   comment.Content,
		CreatedAt: comment.CreatedAt.Format("2006-01-02 15:04:05"),
		TaskId:    comment.TaskId,
		UserId:    comment.UserId,
		User:      rUser,
		Task:      rTask,
	}
	return encodedComment
}

func (this *EncodeCommentService) EncodeMany(comments []*Comment) []*RComment {
	encodedComments := make([]*RComment, 0)
	for _, task := range comments {
		encodedComments = append(encodedComments, this.Encode(task))
	}
	return encodedComments
}

func NewEncodeCommentService(ctx context.Context) *EncodeCommentService {
	service := new(EncodeCommentService)
	service.Ctx = ctx
	return service
}
