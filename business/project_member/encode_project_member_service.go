package project_member
import (
	"context"
	"github.com/kfchen81/beego/vanilla"

)

type EncodeProjectMemberService struct {
	vanilla.ServiceBase
}

//Encode 对单个实体对象进行编码
func (this *EncodeProjectMemberService) Encode(projectMember *ProjectMember) *RProjectMember {
	return &RProjectMember{
		Id: projectMember.Id,
		ProjectId: projectMember.ProjectId,
		UserId: projectMember.UserId,
	}
}

//EncodeMany 对实体对象进行批量编码
func (this *EncodeProjectMemberService) EncodeMany(projectMembers []*ProjectMember) []*RProjectMember {
	rDatas := make([]*RProjectMember, 0)
	for _, projectMember := range projectMembers {
		rDatas = append(rDatas, this.Encode(projectMember))
	}

	return rDatas
}

func NewEncodeProjectMemberService(ctx context.Context) *EncodeProjectMemberService {
	service := new(EncodeProjectMemberService)
	service.Ctx = ctx
	return service
}
