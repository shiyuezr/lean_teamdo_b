package task

import (
	"context"
	"github.com/kfchen81/beego"
	"github.com/kfchen81/beego/vanilla"
	"teamdo/business/account"
	"teamdo/business/lane"
	"teamdo/business/project"
	m_task "teamdo/models/task"
)

type Task struct {
	vanilla.EntityBase
	Id      int
	Content string //任务内容
	Status  int    //任务的完成状态

	UserId       int
	User         *account.User //任务执行者
	ParentTaskId int
	ParentTask   *Task //父级任务
	LaneId       int
	Lane         *lane.Lane //任务所在泳道
	ProjectId    int
	Project      *project.Project //任务所属项目
}

func NewTask(ctx context.Context, content string, status, user_id, parent_task_id, lane_id, project_id int) *Task {
	dbModel := &m_task.Task{
		Content:      content,
		Status:       status,
		UserId:       user_id,
		ParentTaskId: parent_task_id,
		LaneId:       lane_id,
		ProjectId:    project_id,
	}
	_, err := vanilla.GetOrmFromContext(ctx).Insert(dbModel)

	if err != nil {
		beego.Error(err)
		panic(vanilla.NewSystemError("create_task:failed", "创建任务失败"))
	}
	return NewTaskFromDbModel(ctx, dbModel)
}

func NewTaskFromDbModel(ctx context.Context, dbModel *m_task.Task) *Task {
	instance := new(Task)
	instance.Ctx = ctx
	instance.Model = dbModel
	instance.Id = dbModel.Id
	instance.Content = dbModel.Content
	instance.Status = dbModel.Status
	instance.UserId = dbModel.UserId
	instance.ParentTaskId = dbModel.ParentTaskId
	instance.LaneId = dbModel.LaneId
	instance.ProjectId = dbModel.ProjectId
	return instance
}
