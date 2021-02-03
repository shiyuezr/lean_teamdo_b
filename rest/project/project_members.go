package project

import (
	"github.com/kfchen81/beego/vanilla"
	_ "teamdo/business/account"
	"teamdo/business/project"
)

type Members struct {
	vanilla.RestResource
}

func (this *Members) Resource() string {
	return "project.members"
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
	rMembers := project.NewEncodeMemberService(bCtx).EncodeMany(members)

	response := vanilla.MakeResponse(rMembers)
	this.ReturnJSON(response)
}
