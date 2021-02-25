package comment

import (
	"context"
	"github.com/kfchen81/beego"
	"github.com/kfchen81/beego/vanilla"
	"teamdo/business/account"
	"teamdo/business/task"
	m_comment "teamdo/models/comment"
	"time"
)

type Comment struct {
	vanilla.EntityBase
	Id        int
	Content   string
	CreatedAt time.Time
	TaskId    int
	UserId    int
	User      *account.User //评论人
	Task      *task.Task    //所属任务
}

func NewComment(ctx context.Context, content string, taskId, userId int) *Comment {
	dbModel := &m_comment.Comment{
		Content:   content,
		TaskId:    taskId,
		UserId:    userId,
		CreatedAt: time.Now(),
	}
	_, err := vanilla.GetOrmFromContext(ctx).Insert(dbModel)

	if err != nil {
		beego.Error(err)
		panic(vanilla.NewSystemError("create_comment:failed", "创建评论失败"))
	}
	return NewCommentFromDbModel(ctx, dbModel)
}

func NewCommentFromDbModel(ctx context.Context, dbModel *m_comment.Comment) *Comment {
	instance := new(Comment)
	instance.Ctx = ctx
	instance.Model = dbModel
	instance.Id = dbModel.Id
	instance.Content = dbModel.Content
	instance.UserId = dbModel.UserId
	instance.TaskId = dbModel.TaskId
	return instance
}
