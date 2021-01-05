package project

import (
	"github.com/kfchen81/beego/orm"
	"time"
)

type Project struct {
	Id    int
	Name  string

	ManagerId	int
}

func (self *Project) TableName() string {
	return "project_project"
}

type Tunnel struct {
	Id     int
	Title  string
	ProjectId	int
	ManagerId	int

	IsDeleted  bool `orm:"default(false)"`
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
}

func (self *Tunnel) TableName() string {
	return "project_tunnel"
}

type Task struct {
	Id     int
	Title     string
	TunnelId	string
	ExecutorId	int		`orm:"default(0)"`

	Status 		bool	`orm:"default(false)"`
	Remark      string
	Priority    int
	IsDeleted	bool	`orm:"default(false)"`
	Comment  	string	`orm:"size(500);default('')"`
	StartDate	time.Time
	EndDate		time.Time
	CreatedAt	time.Time 	`orm:"auto_now_add;type(datetime)"`
}

func (self *Task) TableName() string {
	return "project_task"
}

// project_has_user
type ProjectHasUser struct {
	Id  		int
	ProjectId	int
	UserId 		int
}

func (self *ProjectHasUser) TableName() string {
	return "project_has_user"
}

func init()  {
	orm.RegisterModel(new(Project))
	orm.RegisterModel(new(Task))
	orm.RegisterModel(new(Tunnel))
	orm.RegisterModel(new(ProjectHasUser))
}
