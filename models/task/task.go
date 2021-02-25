package task

import (
	"github.com/kfchen81/beego/orm"
	"time"
)

type Task struct {
	Id           int
	Content      string    //任务内容
	Status       int       //任务的完成状态
	CreatedAt    time.Time `orm:"auto_now_add;type(datetime)"`
	OperatorId   int       //外键，执行者id
	ParentTaskId int       //外键，父级任务id
	LaneId       int       //外键，任务所在泳道id
	ProjectId    int       //外键，归属项目id
}

func (this *Task) TableName() string {
	return "task_task"
}

type TaskOperationLog struct {
	Id        int
	Content   string
	TaskId    int       //foreign key Task
	UserId    int       //foreign key User
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
}

func (self *TaskOperationLog) TableName() string {
	return "task_operation_log"
}

func init() {
	orm.RegisterModel(new(Task))
	orm.RegisterModel(new(TaskOperationLog))
}
