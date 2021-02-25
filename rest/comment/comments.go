package comment

import (
	"github.com/kfchen81/beego/vanilla"
	b_comment "teamdo/business/comment"
	b_task "teamdo/business/task"
)

type Comments struct {
	vanilla.RestResource
}

func (this *Comments) Resource() string {
	return "comment.comments"
}

func (this *Comments) GetParameters() map[string][]string {
	return map[string][]string{
		"GET": {"task_id:int", "?filters:json", "?with_options:json"},
	}
}

func (this *Comments) Get() {
	taskId, _ := this.GetInt("task_id")
	bCtx := this.GetBusinessContext()
	filters := vanilla.ConvertToBeegoOrmFilter(this.GetFilters())
	task := b_task.NewTaskRepository(bCtx).GetTaskById(taskId)
	if task == nil{
		panic(vanilla.NewBusinessError("task_not_exist", "任务不存在"))
	}

	comments := b_comment.NewCommentRepository(bCtx).GetByTask(task, filters)
	b_comment.NewFillCommentService(bCtx).Fill(comments, this.GetFillOptions("with_options"))
	response := vanilla.MakeResponse(b_comment.NewEncodeCommentService(bCtx).EncodeMany(comments))
	this.ReturnJSON(response)
}