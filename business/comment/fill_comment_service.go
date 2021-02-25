package comment

import (
	"context"
	"github.com/kfchen81/beego/vanilla"
	b_account "teamdo/business/account"
	b_task "teamdo/business/task"
)

type FillCommentService struct {
	vanilla.ServiceBase
}

func (this *FillCommentService) FillOne(comment *Comment, option vanilla.FillOption) {
	this.Fill([]*Comment{comment}, option)
}

func (this *FillCommentService) Fill(comments []*Comment, option vanilla.FillOption) {
	if v, ok := option["with_user"]; ok && v {
		this.fillUser(comments)
	}
	if v, ok := option["with_task"]; ok && v {
		this.fillTask(comments)
	}
}

func (this *FillCommentService) fillUser(comments []*Comment) {
	userIds := make([]int, 0)
	userId2user := make(map[int]struct{})

	for _, task := range comments {
		if _, ok := userId2user[task.UserId]; !ok {
			userIds = append(userIds, task.UserId)
		}
		userId2user[task.UserId] = struct{}{}
	}
	users := b_account.NewUserRepository(this.Ctx).GetByIds(userIds)
	id2user := make(map[int]*b_account.User)
	for _, user := range users{
		id2user[user.Id] = user
	}

	for _, task := range comments{
		task.User = id2user[task.UserId]
	}
}

func (this *FillCommentService) fillTask(comments []*Comment) {
	taskIds := make([]int, 0)
	taskId2task := make(map[int]struct{})

	for _, comment := range comments {
		if _, ok := taskId2task[comment.TaskId]; !ok {
			taskIds = append(taskIds, comment.TaskId)
		}
		taskId2task[comment.TaskId] = struct{}{}
	}
	tasks := b_task.NewTaskRepository(this.Ctx).GetTaskByIds(taskIds)
	id2task := make(map[int]*b_task.Task)
	for _, task := range tasks{
		id2task[task.Id] = task
	}

	for _, comment := range comments{
		comment.Task = id2task[comment.TaskId]
	}
}

func NewFillCommentService(ctx context.Context) *FillCommentService {
	service := new(FillCommentService)
	service.Ctx = ctx
	return service
}
