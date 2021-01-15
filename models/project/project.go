package project

import (
	"github.com/kfchen81/beego/orm"
	"time"
)

type Project struct {
	Id    		int
	Name  		string

	ManagerId	int
	CreateAt	time.Time	`orm:"auto_now_add;type(datetime)"`
	UpdatedAt	time.Time	`orm:"auto_now;type(datetime)"`
}

func (self *Project) TableName() string {
	return "project_project"
}

type Tunnel struct {
	Id     		int
	Title  		string
	ProjectId	int
	ManagerId	int

	IsDeleted  	bool 	`orm:"default(false)"`
	CreatedAt 	time.Time 	`orm:"auto_now_add;type(datetime)"`
	UpdatedAt  	time.Time	`orm:"auto_now;type(datetime)"`
}

func (self *Tunnel) TableName() string {
	return "project_tunnel"
}

// 优先级
const TASK_PRIOTITY_LOWER = 1 // 较低
const TASK_PRIOTITY_ORDINARY = 2 // 普通
const TASK_PRIOTITY_EMERGENCY = 3 // 紧急
const TASK_PRIOTITY_VERY_URGENT = 4 // 非常紧急

var TASK_PRIOTITY_TYPE_CODE2PRIOTITY_TYPE = map[string]int {
	"lower": TASK_PRIOTITY_LOWER,
	"ordinary": TASK_PRIOTITY_ORDINARY,
	"emergency": TASK_PRIOTITY_EMERGENCY,
	"very_urgent": TASK_PRIOTITY_VERY_URGENT,
}

var PRIOTITY_TYPE2CODE = map[int]string {
	TASK_PRIOTITY_LOWER: "lower",
	TASK_PRIOTITY_ORDINARY: "ordinary",
	TASK_PRIOTITY_EMERGENCY: "emergency",
	TASK_PRIOTITY_VERY_URGENT: "very_urgent",
}

type Task struct {
	Id     		int
	Title     	string
	TunnelId	int
	ProjectName string
	ExecutorId	int		`orm:"default(0)"`

	Status 		bool	`orm:"default(false)"`
	Remark      string
	Priority    int
	IsDeleted	bool	`orm:"default(false)"`
	Comment  	string	`orm:"size(500);default('')"`
	StartDate	time.Time
	EndDate		time.Time
	CreatedAt	time.Time 	`orm:"auto_now_add;type(datetime)"`
	UpdatedAt 	time.Time	`orm:"auto_now;type(datetime)"`
}

func (self *Task) TableName() string {
	return "project_task"
}

// project_has_user
type ProjectHasMember struct {
	Id  		int
	ProjectId	int
	UserId 		int
	UpdatedAt	time.Time	`orm:"auto_now;type(datetime)"`
}

func (self *ProjectHasMember) TableName() string {
	return "project_has_user"
}

func init()  {
	orm.RegisterModel(new(Project))
	orm.RegisterModel(new(Task))
	orm.RegisterModel(new(Tunnel))
	orm.RegisterModel(new(ProjectHasMember))
}
