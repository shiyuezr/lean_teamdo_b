package user

import (
	"github.com/kfchen81/beego/vanilla"
	b_user "teamdo/business/user"
)

type Members struct {
	vanilla.RestResource
}

func (this *Members) Resource() string {
	return "member.members"
}

func (this *Members) GetParameters() map[string][]string {
	return map[string][]string{
		"GET": []string{"project_id: int"},
	}
}

func (this *Members) Get()  {
	bCtx := this.GetBusinessContext()
	projectId, _ := this.GetInt("project_id")
	members := b_user.NewMemberRepository(bCtx).GetMembersByProjectId(projectId)
	data := b_user.NewEncodeMemberService(bCtx).EncodeMany(members)
	response := vanilla.MakeResponse(data)
	this.ReturnJSON(response)
}
