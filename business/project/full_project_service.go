package project

import (
	"context"
	b_account "teamdo/business/account"
	m_account "teamdo/models/account"
	m_project "teamdo/models/project"

	"github.com/kfchen81/beego"
	"github.com/kfchen81/beego/vanilla"
)

type FillProjectService struct {
	vanilla.ServiceBase
}

//实例化数据填充服务
func NewFillProjectService(ctx context.Context) *FillProjectService {
	service := new(FillProjectService)
	service.Ctx = ctx
	return service
}

//填充一个
func (this *FillProjectService) FillOne(project *Project, option vanilla.FillOption) {
	this.Fill([]*Project{project}, option)
}

//填充project
func (this *FillProjectService) Fill(projects []*Project, option vanilla.FillOption) {
	if len(projects) == 0 {
		return
	}

	ids := make([]int, 0)
	//将所有的id转成数组类型
	for _, project := range projects {
		ids = append(ids, project.Id)
	}
	/*
		if enableOption, ok := option["with_task"]; ok && enableOption {
			this.fillTask(projects, ids)
		}
	*/
	if enableOption, ok := option["with_user"]; ok && enableOption {
		this.fillUser(projects, ids)
	}

}

/*
这个方法没必要

//1--n关系的数据填充 一个project含有多个task
func (this *FillProjectService) fillTask(projects []*Project, ids []int) {

}


*/

//m-n关系的数据获取方式
func (this *FillProjectService) fillUser(projects []*Project, ids []int) {
	id2entity := make(map[int]*Project)
	//将projects集合拆开
	for _, project := range projects {
		id2entity[project.Id] = project
	}

	o := vanilla.GetOrmFromContext(this.Ctx)

	var relationModels []*m_project.ProjectHasUser
	//使用过滤条件查询project和user的中间表，并把数据放入relationModels中
	_, err := o.QueryTable(&m_project.ProjectHasUser{}).Filter("project_id__in", ids).All(&relationModels)

	if err != nil {
		beego.Error(err)
		return
	}

	if len(relationModels) == 0 {
		return
	}

	userIds := make([]int, 0)

	//将relationModels中的userid的数据解析出来放入userids集合中
	for _, relationModel := range relationModels {
		userIds = append(userIds, relationModel.UserId)
	}

	//使用userid集合从user表中获取user对象，并把数据放入models中
	var models []*m_account.User
	_, err = o.QueryTable(&m_account.User{}).Filter("id__in", userIds).All(&models)

	if err != nil {
		beego.Error(err)
		return
	}

	id2model := make(map[int]*m_account.User)

	//将模型拆开放入集合
	for _, model := range models {
		id2model[model.Id] = model
	}
	//将user数据整合进入project
	for _, relationModel := range relationModels {
		projectId := relationModel.ProjectId
		userId := relationModel.UserId

		if project, ok := id2entity[projectId]; ok {
			if model, ok2 := id2model[userId]; ok2 {
				project.Users = append(project.Users, b_account.NewUserFromModel(this.Ctx, model))
			}
		}
	}
}
