package lane

import (
	"context"
	"github.com/kfchen81/beego"
	"github.com/kfchen81/beego/vanilla"
	models "teamdo/models/lane"
)


type LaneRepository struct {
	vanilla.RepositoryBase
}

func (this *LaneRepository) GetLanes(filters vanilla.Map, orderExprs ...string) []*Lane {
	o := vanilla.GetOrmFromContext(this.Ctx)
	qs := o.QueryTable(&models.Lane{})

	var models []*models.Lane
	if len(filters) > 0 {
		qs = qs.Filter(filters)
	}
	if len(orderExprs) > 0 {
		qs = qs.OrderBy(orderExprs...)
	}
	_, err := qs.All(&models)
	if err != nil {
		beego.Error(err)
		return nil
	}

	lanes := make([]*Lane, 0)
	for _, model := range models {
		lanes = append(lanes, NewLaneFromDbModel(this.Ctx, model))
	}
	return lanes
}


func (this *LaneRepository) GetLane(id int) *Lane {
	filters := vanilla.Map{
		"id": id,
	}

	lanes := this.GetLanes(filters)

	if len(lanes) == 0 {
		return nil
	} else {
		return lanes[0]
	}
}

func NewLaneRepository(ctx context.Context) *LaneRepository {
	inst := new(LaneRepository)
	inst.Ctx = ctx
	return inst
}