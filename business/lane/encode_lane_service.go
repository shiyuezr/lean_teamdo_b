package lane

import (
	"context"
	"github.com/kfchen81/beego/vanilla"
)

type EncodeLaneService struct {
	vanilla.ServiceBase
}

// Encode 对单个实体对象进行编码
func (this *EncodeLaneService) Encode(lane *Lane) *RLane {
	if lane == nil {
		return nil
	}

	encodedLane := &RLane{
		Id:   lane.Id,
		Name: lane.Name,
	}
	return encodedLane
}

func (this *EncodeLaneService) EncodeMany(lanes []*Lane) []*RLane {
	encodedStores := make([]*RLane, 0)
	for _, lane := range lanes {
		encodedStores = append(encodedStores, this.Encode(lane))
	}
	return encodedStores
}

func NewEncodeLaneService(ctx context.Context) *EncodeLaneService {
	service := new(EncodeLaneService)
	service.Ctx = ctx
	return service
}
