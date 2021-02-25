package lane

import "github.com/kfchen81/beego/orm"

type Lane struct {
	Id        int
	Name      string
	Sort      string
	ProjectId int //foreign key Project
}

func (this *Lane) TableName() string {
	return "lane_lane"
}

func init() {
	orm.RegisterModel(new(Lane))
}
