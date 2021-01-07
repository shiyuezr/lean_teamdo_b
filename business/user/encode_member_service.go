package user

import (
	"context"
	"github.com/kfchen81/beego/vanilla"
)

type EncodeMemberService struct {
	vanilla.ServiceBase
}

func NewEncodeMemberService(ctx context.Context) *EncodeMemberService {
	service := new(EncodeMemberService)
	service.Ctx = ctx
	return service
}

func (this *EncodeMemberService) Encode(member *ProjectMember) *RMember {
	if member == nil {
		return nil
	}

	return &RMember{
		Name: member.UserName,
	}
}

func (this *EncodeMemberService) EncodeMany(members []*ProjectMember) []*RMember {
	rDatas := make([]*RMember, 0)
	for _, member := range members {
		rDatas = append(rDatas, this.Encode(member))
	}
	return rDatas
}