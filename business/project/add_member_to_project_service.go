package project

import (
	"context"
	"github.com/kfchen81/beego"
	"github.com/kfchen81/beego/vanilla"
	m_project "teamdo/models/project"
	"time"
)

type AddMemberToProjectService struct {
	vanilla.ServiceBase
}

func (this *AddMemberToProjectService) AddMemberToProject(projectId, userId int)  {
	o := vanilla.GetOrmFromContext(this.Ctx)
	nowTime := time.Now()
	var err error
	qs := o.QueryTable(m_project.ProjectHasMember{}).Filter(vanilla.Map{
		"project_id": projectId,
		"user_id": userId,
	})

	if !qs.Exist() {
		_, err = o.Insert(&m_project.ProjectHasMember{
			ProjectId: projectId,
			UserId: userId,
			UpdatedAt: nowTime,
		})
	}

	if err != nil {
		beego.Error(err)
		panic(vanilla.NewSystemError("update_project_member:failed", "更新项目成员失败"))
	}
}

func NewProjectMemberService(ctx context.Context) *AddMemberToProjectService {
	instance := new(AddMemberToProjectService)
	instance.Ctx = ctx
	return instance
}