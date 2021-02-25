package task_operation_log

import (
	"context"
	"github.com/kfchen81/beego"
	"github.com/kfchen81/beego/vanilla"
	"teamdo/business/account"
	"teamdo/business/task"
	m_task "teamdo/models/task"
	"time"
)

type TaskOperationLog struct {
	vanilla.EntityBase
	Id        int
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
	Content   string
	TaskId    int
	UserId    int

	Task *task.Task    //日志所属任务
	User *account.User //日志撰写者
}

func NewTaskOperationLog(ctx context.Context, content string, taskId, userId int) *TaskOperationLog {
	dbModel := &m_task.TaskOperationLog{
		Content:   content,
		UserId:    userId,
		TaskId:    taskId,
		CreatedAt: time.Now(),
	}
	_, err := vanilla.GetOrmFromContext(ctx).Insert(dbModel)

	if err != nil {
		beego.Error(err)
		panic(vanilla.NewSystemError("create_task_operation_log:failed", "创建变动日志失败"))
	}
	return NewTaskOperationLogFromDbModel(ctx, dbModel)
}

func NewTaskOperationLogFromDbModel(ctx context.Context, dbModel *m_task.TaskOperationLog) *TaskOperationLog {
	instance := new(TaskOperationLog)
	instance.Ctx = ctx
	instance.Model = dbModel
	instance.Id = dbModel.Id
	instance.Content = dbModel.Content
	instance.CreatedAt = dbModel.CreatedAt
	instance.UserId = dbModel.UserId
	instance.TaskId = dbModel.TaskId
	return instance
}
