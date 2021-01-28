package project

import (
	"context"
	"errors"
	"github.com/kfchen81/beego"
	"github.com/kfchen81/beego/orm"
	"github.com/kfchen81/beego/vanilla"
	"teamdo/business/account"
	m_project "teamdo/models/project"
)

type Project struct {
	vanilla.EntityBase
	Id      int
	Name    string
	Content string
	Status  int //完成状态

	Administrators []*account.User //管理员
	Participants   []*account.User //参与者
}

func (this *Project) Update(name string, content string, status int) error {
	var model m_project.Project
	o := vanilla.GetOrmFromContext(this.Ctx)
	_, err := o.QueryTable(&model).Filter("id", this.Id).Update(orm.Params{
		"name":    name,
		"content": content,
		"status":  status,
	})
	if err != nil {
		beego.Error(err)
		return errors.New("project:update_fail")
	}
	return nil

}

func (this *Project) Delete() {
	_, err := vanilla.GetOrmFromContext(this.Ctx).QueryTable(m_project.Project{}).Filter(vanilla.Map{"id": this.Id}).Delete()
	if err != nil {
		beego.Error(err)
		panic(err)
	}
}

func NewProject(ctx context.Context, name string, content string) *Project {
	o := vanilla.GetOrmFromContext(ctx)
	qs := o.QueryTable(&m_project.Project{})

	if qs.Filter("name", name).Exist() {
		panic(vanilla.NewBusinessError("project_name:existed", "项目名已存在"))
	}
	dbModel := &m_project.Project{
		Name:    name,
		Content: content,
		Status:  1,
	}
	_, err := o.Insert(dbModel)
	if err != nil {
		beego.Error(err)
		panic(vanilla.NewSystemError("create_project:failed", "创建项目失败"))
	}
	return NewProjectFromDbModel(ctx, dbModel)
}

func NewProjectFromDbModel(ctx context.Context, dbModel *m_project.Project) *Project {
	instance := new(Project)
	instance.Ctx = ctx
	instance.Model = dbModel
	instance.Id = dbModel.Id
	instance.Name = dbModel.Name
	instance.Content = dbModel.Content
	return instance
}
