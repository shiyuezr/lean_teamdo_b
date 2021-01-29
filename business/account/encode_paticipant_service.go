package account

import (
	"context"
	"github.com/kfchen81/beego/vanilla"
)

type EncodeParticipantService struct {
	vanilla.ServiceBase
}

func NewEncodeParticipantService(ctx context.Context) *EncodeParticipantService {
	service := new(EncodeParticipantService)
	service.Ctx = ctx
	return service
}

//Encode 对单个实体对象进行编码
func (this *EncodeParticipantService) Encode(participant *User) *RUser {
	if participant == nil {
		return nil
	}

	return &RUser{
		Id:       participant.Id,
		Username: participant.Username,
	}
}

//EncodeMany 对实体对象进行批量编码
func (this *EncodeParticipantService) EncodeMany(participants []*User) []*RUser {
	rDatas := make([]*RUser, 0)
	for _, participant := range participants {
		rDatas = append(rDatas, this.Encode(participant))
	}
	return rDatas
}

func init() {
}
