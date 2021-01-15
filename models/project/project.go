package project

import (
	"time"

	"github.com/kfchen81/beego/orm"
)

//
type State int

const (
	Running State = iota
	Pending
	Proceed
	Done
)

//定义表的字段
type Project struct {
	Id           int
	Name         string
	Introduction string
	Cover        string
	StartTime    time.Time
	CreateAt     time.Time `orm:"auto_now_add;type(datetime)"`
	IsEnabled    bool
	IsDeleted    bool
}

//定义表名
func (this *Project) TableName() string {
	return "project_project"
}

/*
//创建数据库的索引
func (this *Project) TableIndex() [][]string {
	return [][]string{
		[]string{"IsDeleted", "IsEnabled"},
	}
}
*/

type ProjectHasUser struct {
	Id        int
	ProjectId int
	UserId    int
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
}

//定义表名
func (this *ProjectHasUser) TableName() string {
	return "project_project_has_user"
}

//Task
type Level int
type Task struct {
	Id        int
	Title     string
	Status    int //1:待处理 2:进行中 3:已完成
	IsEnabled bool
	IsDeleted bool
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
	Deadline  time.Time
	Priority  int //1.low 2.nomal 3.urgent 4.very_urgent
}

//定义表名
func (this *Task) TableName() string {
	return "project_task"
}

//暂时还没搞懂什么意思
func (this *Task) TableIndex() [][]string {
	return [][]string{
		[]string{"IsDeleted", "IsEnabled"},
	}
}

type TaskHasUser struct {
	Id        int
	TaskId    int
	UserId    int
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
}

//定义表名
func (this *TaskHasUser) TableName() string {
	return "project_task_has_user"
}

type TaskLog struct {
	Id        int
	TaskId    int
	Content   string
	IsEnabled bool
	IsDeleted bool
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
	Deadline  time.Time
	Priority  int //1.low 2.nomal 3.urgent 4.very_urgent
}

//定义表名
func (this *TaskLog) TableName() string {
	return "project_tasklog"
}

/*
//暂时还没搞懂什么意思
func (this *TaskLog) TableIndex() [][]string {
	return [][]string{
		[]string{"IsDeleted", "IsEnabled"},
	}
}
*/
//Comment 评论表
type Comment struct {
	Id        int
	UserId    int
	Content   string
	IsEnabled bool
	IsDeleted bool
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
}

//TableName 定义表名
func (this *Comment) TableName() string {
	return "project_comment"
}

/*
func (this *Comment) TableIndex() [][]string {
	return [][]string{
		[]string{"IsDeleted", "IsEnabled"},
	}
}
*/

func init() {
	orm.RegisterModel(new(Project))
	orm.RegisterModel(new(Task))
	orm.RegisterModel(new(TaskHasUser))
	orm.RegisterModel(new(TaskLog))
	orm.RegisterModel(new(Comment))
}
