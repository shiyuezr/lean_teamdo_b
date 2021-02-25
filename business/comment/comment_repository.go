package comment

import (
	"context"
	"github.com/kfchen81/beego"
	"github.com/kfchen81/beego/vanilla"
	b_task "teamdo/business/task"
	m_comment "teamdo/models/comment"
)

type CommentRepository struct {
	vanilla.RepositoryBase
}

func (this *CommentRepository) GetByFilters(filters vanilla.Map) []*Comment {
	qs := vanilla.GetOrmFromContext(this.Ctx).QueryTable(&m_comment.Comment{})
	if len(filters) > 0 {
		qs = qs.Filter(filters)
	}

	var dbModels []*m_comment.Comment
	_, err := qs.OrderBy("-id").All(&dbModels)
	if err != nil {
		beego.Error(err)
		return []*Comment{}
	}
	comments := make([]*Comment, 0, len(dbModels))
	for _, dbModel := range dbModels {
		comments = append(comments, NewCommentFromDbModel(this.Ctx, dbModel))
	}
	return comments
}

func (this *CommentRepository) GetCommentById(id int) *Comment {
	filters := vanilla.Map{
		"id": id,
	}
	comments := this.GetByFilters(filters)
	if len(comments) == 0 {
		return nil
	}
	return comments[0]
}

func(this *CommentRepository) GetCommentByIds(ids []int) []*Comment {
	filters := vanilla.Map{
		"id__in": ids,
	}
	comments := this.GetByFilters(filters)
	if len(comments) == 0 {
		return nil
	}
	return comments
}

func (this *CommentRepository) GetByTask(task *b_task.Task, filters vanilla.Map) []*Comment{
	filters["task_id"] = task.Id
	return this.GetByFilters(filters)
}

func NewCommentRepository(ctx context.Context) *CommentRepository {
	repository := new(CommentRepository)
	repository.Ctx = ctx
	return repository
}
