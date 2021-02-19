package task

import (
"context"
"github.com/kfchen81/beego"
"github.com/kfchen81/beego/vanilla"
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

func NewTaskRepository(ctx context.Context) *TaskRepository {
	repository := new(TaskRepository)
	repository.Ctx = ctx
	return repository
}
