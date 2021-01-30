package project

import (
	"github.com/kfchen81/beego/vanilla"
	b_project "teamdo/business/project"
)

type ProjectManager struct {
	vanilla.RestResource
}

func (this *ProjectManager) Resource() string {
	return "project.manager"
}

func (this *ProjectManager) GetParameters() map[string][]string {
	return map[string][]string{
		"PUT":    []string{"uid:int", "pid:int"},
		"DELETE": []string{"id:int"},
	}
}

func (this *ProjectManager) Put() {
	uid, _ := this.GetInt("uid")
	pid, _ := this.GetInt("pid")
	bCtx := this.GetBusinessContext()

	//缺少权限验证，是管理员才能添加
	//获取当前用户uid从new，fromcontext //好像再验证方法里，不用写
	//有当前项目id，然后去仓库获取对象，调用验证方法

	//要从仓库中创建project对象，而不是这样直接调用
	//好像也可以
	//参考sparrow的product的put的new
	p_a:= b_project.NewProjectToAdministrator(bCtx, pid, uid)

	response := vanilla.MakeResponse(vanilla.Map{
		"id": p_a.Id,
	})
	this.ReturnJSON(response)
}
