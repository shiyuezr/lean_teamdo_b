package lane

import (
	"context"
	"github.com/kfchen81/beego"
	"github.com/kfchen81/beego/vanilla"
	m_lane "teamdo/models/lane"
)

type LaneRepository struct {
	vanilla.RepositoryBase
}

func (this *LaneRepository) GetByFilters(filters vanilla.Map) []*Lane {
	qs := vanilla.GetOrmFromContext(this.Ctx).QueryTable(&m_lane.Lane{})
	if len(filters) > 0 {
		qs = qs.Filter(filters)
	}

	var dbModels []*m_lane.Lane
	_, err := qs.OrderBy("-id").All(&dbModels)
	if err != nil {
		beego.Error(err)
		return []*Lane{}
	}
	lanes := make([]*Lane, 0, len(dbModels))
	for _, dbModel := range dbModels {
		lanes = append(lanes, NewLaneFromDbModel(this.Ctx, dbModel))
	}
	return lanes
}

func (this *LaneRepository) GetLaneById(id int) *Lane {
	filters := vanilla.Map{
		"id": id,
	}
	lanes := this.GetByFilters(filters)
	if len(lanes) == 0 {
		return nil
	}
	return lanes[0]
}

func NewLaneRepository(ctx context.Context) *LaneRepository {
	repository := new(LaneRepository)
	repository.Ctx = ctx
	return repository
}
