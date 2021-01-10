package project

import (
	"github.com/kfchen81/beego/vanilla"
	"teamdo/business/project"
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
	members := project.NewMemberRepository(bCtx).GetMembersByProjectId(projectId)
	data := project.NewEncodeMemberService(bCtx).EncodeMany(members)
	response := vanilla.MakeResponse(data)
	this.ReturnJSON(response)
}
