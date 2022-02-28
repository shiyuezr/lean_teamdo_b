package lane

import "github.com/kfchen81/beego/orm"

type Lane struct {
	Id int
	Name string
	ProjectId int
	SortId int
	IsDelete bool
}

func (this *Lane) TableName()  string {
	return "lane_lane"
}
func init()  {
	orm.RegisterModel(new(Lane))
}