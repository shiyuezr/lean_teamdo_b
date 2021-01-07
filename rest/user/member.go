package user

import (
	"github.com/kfchen81/beego/vanilla"
	b_project "teamdo/business/project"
)

type Member struct {
	vanilla.RestResource
}

func (this *Member) Resource() string {
	return "member.member"
}

func (this *Member) GetParameter() map[string][]string {
	return map[string][]string{
		"POST": []string{"project_id: int", "user_id: int"},
	}
}

func (this *Member) Post()  {
	bCtx := this.GetBusinessContext()
	projectId, _ := this.GetInt("project_id")
	userId, _ := this.GetInt("user_id")
	b_project.NewProjectMemberService(bCtx).UpdateProjectToMember(projectId, userId)
	response := vanilla.MakeResponse(vanilla.Map{})
	this.ReturnJSON(response)
}
