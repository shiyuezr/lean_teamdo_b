package account

import (
	"context"
	"github.com/kfchen81/beego/vanilla"
)

type EncodeAdministratorService struct {
	vanilla.ServiceBase
}

func NewEncodeAdministratorService(ctx context.Context) *EncodeAdministratorService {
	service := new(EncodeAdministratorService)
	service.Ctx = ctx
	return service
}

// Encode 对单个实体对象进行编码
func (this *EncodeAdministratorService) Encode(administrator *User) *RUser {
	if administrator == nil {
		return nil
	}

	return &RUser{
		Id:       administrator.Id,
		Username: administrator.Username,
	}
}

// EncodeMany 对实体对象进行批量编码
func (this *EncodeAdministratorService) EncodeMany(administrators []*User) []*RUser {
	rDatas := make([]*RUser, 0)
	for _, administrator := range administrators {
		rDatas = append(rDatas, this.Encode(administrator))
	}

	return rDatas
}

func init() {
}
