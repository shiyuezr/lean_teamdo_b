package project

import (
	"context"
	"github.com/kfchen81/beego/vanilla"
	b_user "teamdo/business/user"
)

type EncodeMemberService struct {
	vanilla.ServiceBase
}

func NewEncodeMemberService(ctx context.Context) *EncodeMemberService {
	service := new(EncodeMemberService)
	service.Ctx = ctx
	return service
}

func (this *EncodeMemberService) Encode(member *ProjectMember) *b_user.RMember {
	if member == nil {
		return nil
	}

	return &b_user.RMember{
		Id: member.Id,
		Name: member.UserName,
	}
}

func (this *EncodeMemberService) EncodeMany(members []*ProjectMember) []*b_user.RMember {
	rDatas := make([]*b_user.RMember, 0)
	for _, member := range members {
		rDatas = append(rDatas, this.Encode(member))
	}
	return rDatas
}