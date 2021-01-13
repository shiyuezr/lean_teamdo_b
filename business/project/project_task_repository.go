package project

import (
	"context"
	"github.com/kfchen81/beego"
	"github.com/kfchen81/beego/vanilla"
	m_project "teamdo/models/project"
)

type TaskRepository struct {
	vanilla.RepositoryBase
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

func (this *TaskRepository) GetTasksByTunnelIds(tunnelIds []int) []*Task {
	if len(tunnelIds) == 0 {
		return nil
	}
	filters := vanilla.Map{
		"tunnel_id__in": tunnelIds,
		"is_delete": false,
	}
	tasks := this.GetByFilters(filters)
	return tasks
}

func (this *TaskRepository) GetByFilters(filters vanilla.Map) []*Task {
	o := vanilla.GetOrmFromContext(this.Ctx)
	qs := o.QueryTable(&m_project.Project{})

	var models  []*m_project.Task
	if len(filters) > 0 {
		qs = qs.Filter(filters)
	}

	_, err := qs.All(&models)
	if err != nil {
		beego.Error(err)
		return nil
	}
	tasks := make([]*Task, 0)
	for _, model := range models {
		tasks = append(tasks, NewTaskForModel(this.Ctx, model))
	}
	return tasks
}

func NewTaskRepository(ctx context.Context) *TaskRepository {
	repository := new(TaskRepository)
	repository.Ctx = ctx
	return repository
}

