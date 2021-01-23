package task

import (
	"github.com/kfchen81/beego/vanilla"
	"teamdo/business/comment"
	"teamdo/business/log"
)

type Task struct {
	vanilla.EntityBase
	Id          int
	Content     string //任务内容
	ProjectTask string //所属项目
	UserTask    string //任务执行者名字
	Status      string //任务的完成状态

	SubTask  []*Task            //子任务
	Comments []*comment.Comment //任务的评论
	LogTask  []*log.Log         //任务的日志
}
