package task

import (
	"github.com/kfchen81/beego/vanilla"
	b_account "teamdo/business/account"
	b_lane "teamdo/business/lane"
	b_project "teamdo/business/project"
	b_task "teamdo/business/task"
)

type Tasks struct {
	vanilla.RestResource
}

func (this *Tasks) Resource() string {
	return "task.tasks"
}

func (this *Tasks) GetParameters() map[string][]string {
	return map[string][]string{
		"GET": {"product_id:int","?lane_id:int", "?operator_id:int", "?filters:json", "?with_options:json"},
	}
}

func (this *Tasks) Get() {
	productId, _ := this.GetInt("product_id")
	operatorId, _ := this.GetInt("operator_id")
	laneId, _ := this.GetInt("lane_id")
	bCtx := this.GetBusinessContext()
	filters := vanilla.ConvertToBeegoOrmFilter(this.GetFilters())
	project := b_project.NewProjectRepository(bCtx).GetProjectById(productId)
	if project == nil{
		panic(vanilla.NewBusinessError("product_not_exist", "项目不存在"))
	}
	operator := b_account.NewUserRepository(bCtx).GetById(operatorId)
	lane := b_lane.NewLaneRepository(bCtx).GetLaneById(laneId)
	tasks := b_task.NewTaskRepository(bCtx).GetByProductAndOperator(project, operator, lane, filters)

	b_task.NewFillTaskService(bCtx).Fill(tasks, this.GetFillOptions("with_options"))
	response := vanilla.MakeResponse(b_task.NewEncodeTaskService(bCtx).EncodeMany(tasks))
	this.ReturnJSON(response)
}