package comment

import (
	//"github.com/kfchen81/beego"
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

//func (this *CommentRepository) GetComments(filters vanilla.Map, orderExprs ...string) []*Comment {
//	o := vanilla.GetOrmFromContext(this.Ctx)
//	qs := o.QueryTable(&Comment{})
//
//	var models []*Comment
//	if len(filters) > 0 {
//		qs = qs.Filter(filters)
//	}
//	if len(orderExprs) > 0 {
//		qs = qs.OrderBy(orderExprs...)
//	}
//	_, err := qs.All(&models)
//	if err != nil {
//		beego.Error(err)
//		return nil
//	}
//
//	comments := make([]*Comment, 0)
//	for _, model := range models {
//		commments = append(categories, NewCommentFromModel(this.Ctx, model))
//	}
//	return comments
//}

////GetComment 根据id和corp获得Comment对象
//func (this *CommentRepository) GetComment(id int) *Comment {
//	filters := vanilla.Map{
//		"id": id,
//	}
//
//	Comments := this.GetComments(filters)
//
//	if len(Comments) == 0 {
//		return nil
//	} else {
//		return Comments[0]
//	}
//}
