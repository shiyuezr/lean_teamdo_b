package project

import (
	"context"
	"github.com/kfchen81/beego/vanilla"
	"teamdo/business/user"
)

type EncodeMemberService struct {
	vanilla.ServiceBase
}

func NewEncodeMemberService(ctx context.Context) *EncodeMemberService {
	service := new(EncodeMemberService)
	service.Ctx = ctx
	return service
}

func (this *EncodeMemberService) Encode(member *ProjectMember) *user.RMember {
	if member == nil {
		return nil
	}

	return &user.RMember{
		Name: member.UserName,
	}
}

func (this *EncodeMemberService) EncodeMany(members []*ProjectMember) []*user.RMember {
	rDatas := make([]*user.RMember, 0)
	for _, member := range members {
		rDatas = append(rDatas, this.Encode(member))
	}
	return rDatas
}