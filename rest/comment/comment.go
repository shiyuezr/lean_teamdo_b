package comment

import (
	"github.com/kfchen81/beego/vanilla"
	b_account "teamdo/business/account"
	b_comment "teamdo/business/comment"
)

type Comment struct {
	vanilla.RestResource
}

func (this *Comment) Resource() string {
	return "comment.comment"
}

func (this *Comment) GetParameters() map[string][]string {
	return map[string][]string{
		"GET": []string{"id:int", "?with_options:json"},
		"PUT": []string{
			"content:string",
			"task_id:int",
		},
	}
}

func (this *Comment) Put() {
	content := this.GetString("content")
	taskId, _ := this.GetInt("task_id")
	bCtx := this.GetBusinessContext()
	userId := b_account.GetUserFromContext(bCtx).Id

	comment := b_comment.NewComment(bCtx, content, taskId, userId)
	response := vanilla.MakeResponse(vanilla.Map{
		"id": comment.Id,
	})
	this.ReturnJSON(response)
}

func (this *Comment) Get() {
	id, _ := this.GetInt("id")
	bCtx := this.GetBusinessContext()

	comment := b_comment.NewCommentRepository(bCtx).GetCommentById(id)
	b_comment.NewFillCommentService(bCtx).FillOne(comment, this.GetFillOptions("with_options"))

	encodeService := b_comment.NewEncodeCommentService(bCtx)
	respData := encodeService.Encode(comment)
	response := vanilla.MakeResponse(respData)
	this.ReturnJSON(response)
}

