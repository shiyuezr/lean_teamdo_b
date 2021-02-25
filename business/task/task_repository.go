package task

import (
	"context"
	"github.com/kfchen81/beego"
	"github.com/kfchen81/beego/vanilla"
	b_account "teamdo/business/account"
	b_lane "teamdo/business/lane"
	b_project "teamdo/business/project"
	m_task "teamdo/models/task"
)

type TaskRepository struct {
	vanilla.RepositoryBase
}

func (this *TaskRepository) GetByFilters(filters vanilla.Map) []*Task {
	qs := vanilla.GetOrmFromContext(this.Ctx).QueryTable(&m_task.Task{})
	if len(filters) > 0 {
		qs = qs.Filter(filters)
	}

	var dbModels []*m_task.Task
	_, err := qs.OrderBy("-id").All(&dbModels)
	if err != nil {
		beego.Error(err)
		return []*Task{}
	}
	tasks := make([]*Task, 0, len(dbModels))
	for _, dbModel := range dbModels {
		tasks = append(tasks, NewTaskFromDbModel(this.Ctx, dbModel))
	}
	return tasks
}

func (this *TaskRepository) GetTaskById(id int) *Task {
	filters := vanilla.Map{
		"id": id,
	}
	tasks := this.GetByFilters(filters)
	if len(tasks) == 0 {
		return nil
	}
	return tasks[0]
}

func(this *TaskRepository) GetTaskByIds(ids []int) []*Task {
	filters := vanilla.Map{
		"id__in": ids,
	}
	tasks := this.GetByFilters(filters)
	if len(tasks) == 0 {
		return nil
	}
	return tasks
}

func (this *TaskRepository) GetByProductAndOperator(project *b_project.Project, operator *b_account.User, lane *b_lane.Lane, filters vanilla.Map) []*Task{
	if operator == nil {
		filters["project_id"] = project.Id
		filters["lane_id"] = lane.Id
		return this.GetByFilters(filters)
	} else {
		filters["project_id"] = project.Id
		filters["operator_id"] = operator.Id
		return this.GetByFilters(filters)
	}
}

func NewTaskRepository(ctx context.Context) *TaskRepository {
	repository := new(TaskRepository)
	repository.Ctx = ctx
	return repository
}
