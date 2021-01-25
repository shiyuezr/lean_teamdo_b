package comment

import (
	"context"
	"github.com/kfchen81/beego/vanilla"
)

type CommentRepository struct {
	vanilla.RepositoryBase
}

func NewCommentRepository(ctx context.Context) *CommentRepository {
	repository := new(CommentRepository)
	repository.Ctx = ctx
	return repository
}
