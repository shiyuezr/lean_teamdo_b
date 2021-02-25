package task

import (
	"context"
	"github.com/kfchen81/beego/vanilla"
	b_account "teamdo/business/account"
	b_lane "teamdo/business/lane"
	b_project "teamdo/business/project"
)

type FillTaskService struct {
	vanilla.ServiceBase
}

func (this *FillTaskService) FillOne(task *Task, option vanilla.FillOption) {
	this.Fill([]*Task{task}, option)
}

func (this *FillTaskService) Fill(tasks []*Task, option vanilla.FillOption) {
	if v, ok := option["with_operator"]; ok && v {
		this.fillOperator(tasks)
	}
	if v, ok := option["with_parent_task"]; ok && v {
		//this.fillParentTask(tasks)
	}
	if v, ok := option["with_lane"]; ok && v {
		this.fillLane(tasks)
	}
	if v, ok := option["with_project"]; ok && v {
		this.fillProject(tasks)
	}
}

func (this *FillTaskService) fillOperator(tasks []*Task) {
	operatorIds := make([]int, 0)
	operatorId2operator := make(map[int]struct{})

	for _, task := range tasks {
		if _, ok := operatorId2operator[task.OperatorId]; !ok {
			operatorIds = append(operatorIds, task.OperatorId)
		}
		operatorId2operator[task.OperatorId] = struct{}{}
	}
	operators := b_account.NewUserRepository(this.Ctx).GetByIds(operatorIds)
	id2operator := make(map[int]*b_account.User)
	for _, operator := range operators{
		id2operator[operator.Id] = operator
	}

	for _, task := range tasks{
		task.Operator = id2operator[task.OperatorId]
	}
}

func (this *FillTaskService) fillLane(tasks []*Task) {
	laneIds := make([]int, 0)
	laneId2lane := make(map[int]struct{})

	for _, task := range tasks {
		if _, ok := laneId2lane[task.LaneId]; !ok {
			laneIds = append(laneIds, task.LaneId)
		}
		laneId2lane[task.LaneId] = struct{}{}
	}
	lanes := b_lane.NewLaneRepository(this.Ctx).GetLanesByIds(laneIds)
	id2lane := make(map[int]*b_lane.Lane)
	for _, lane := range lanes{
		id2lane[lane.Id] = lane
	}

	for _, task := range tasks{
		task.Lane = id2lane[task.LaneId]
	}
}

func (this *FillTaskService) fillProject(tasks []*Task) {
	projectIds := make([]int, 0)
	projectId2project := make(map[int]struct{})

	for _, task := range tasks {
		if _, ok := projectId2project[task.ProjectId]; !ok {
			projectIds = append(projectIds, task.ProjectId)
		}
		projectId2project[task.ProjectId] = struct{}{}
	}
	projects := b_project.NewProjectRepository(this.Ctx).GetProjectByIds(projectIds)
	id2project := make(map[int]*b_project.Project)
	for _, operator := range projects{
		id2project[operator.Id] = operator
	}

	for _, task := range tasks{
		task.Project = id2project[task.ProjectId]
	}
}

func NewFillTaskService(ctx context.Context) *FillTaskService {
	inst := new(FillTaskService)
	inst.Ctx = ctx
	return inst
}
