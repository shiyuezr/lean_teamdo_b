package project

import "github.com/kfchen81/beego/vanilla"

type Task struct {
	vanilla.EntityBase
	Id 			int
	Title       string
	TunnelId	int
	ExecutorId  int

	Status 		bool
	IsDelete	bool
	Remark		string
	Priority	string
}
