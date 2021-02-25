package task

import (
	"context"
	"fmt"
	"github.com/kfchen81/beego"
	"github.com/kfchen81/beego/orm"
	"github.com/kfchen81/beego/vanilla"
	"teamdo/business/account"
	"teamdo/business/lane"
	"teamdo/business/project"
	m_task "teamdo/models/task"
	"time"
)

type Task struct {
	vanilla.EntityBase
	Id        int
	Content   string    //任务内容
	Status    int       //任务的完成状态
	CreatedAt time.Time //任务的创建时间

	OperatorId   int
	Operator     *account.User //任务执行者
	ParentTaskId int
	ParentTask   *Task //父级任务
	LaneId       int
	Lane         *lane.Lane //任务所在泳道
	ProjectId    int
	Project      *project.Project //任务所属项目
}

func (this *Task) Update(status int, content string, laneId, operatorId int) {
	var model m_task.Task
	o := vanilla.GetOrmFromContext(this.Ctx)
	_, err := o.QueryTable(&model).Filter("id", this.Id).Update(orm.Params{
		"status": status,
		"content": content,
		"operator_id": operatorId,
		"lane_id": laneId,
	})
	if err != nil {
		beego.Error(err)
		panic(vanilla.NewBusinessError("update：fail", fmt.Sprintf("修改任务失败")))
	}
}

func (this *Task) Delete() {
	_, err := vanilla.GetOrmFromContext(this.Ctx).QueryTable(m_task.Task{}).Filter(vanilla.Map{"id": this.Id}).Delete()
	if err != nil {
		beego.Error(err)
		panic(err)
	}
}

func NewTask(ctx context.Context, content string, operatorId, parentTaskId, laneId, projectId int) *Task {
	dbModel := &m_task.Task{
		Content:      content,
		Status:       0, //任务创建时默认为未完成状态
		OperatorId:   operatorId,
		ParentTaskId: parentTaskId,
		LaneId:       laneId,
		ProjectId:    projectId,
		CreatedAt:    time.Now(),
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
	instance.OperatorId = dbModel.OperatorId
	instance.ParentTaskId = dbModel.ParentTaskId
	instance.LaneId = dbModel.LaneId
	instance.ProjectId = dbModel.ProjectId
	return instance
}
