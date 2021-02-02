package lane

import (
	"github.com/kfchen81/beego/vanilla"
	b_lane "teamdo/business/lane"
)

type Lanes struct {
	vanilla.RestResource
}

// Resource 项目列表
func (this *Lanes) Resource() string {
	return "lane.lanes"
}

func (this *Lanes) GetParameters() map[string][]string {
	return map[string][]string{
		"GET": {"?filters:json"},
	}
}

func (this *Lanes) Get() {
	bCtx := this.GetBusinessContext()
	filters := vanilla.ConvertToBeegoOrmFilter(this.GetFilters())

	lanes := b_lane.NewLaneRepository(bCtx).GetByFilters(filters)
	response := vanilla.MakeResponse(b_lane.NewEncodeLaneService(bCtx).EncodeMany(lanes))
	this.ReturnJSON(response)
}
