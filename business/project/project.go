package project

import (
	"context"
	"fmt"
	"github.com/kfchen81/beego"
	"github.com/kfchen81/beego/orm"
	"github.com/kfchen81/beego/vanilla"
	b_account "teamdo/business/account"
	m_project "teamdo/models/project"
)

type Project struct {
	vanilla.EntityBase
	Id      int
	Name    string
	Content string
	Status  int //完成状态

	Administrators []*b_account.User //管理员
	Participants   []*b_account.User //参与者
}

func (this *Project) Update(name string, content string, status int) {
	var model m_project.Project
	o := vanilla.GetOrmFromContext(this.Ctx)
	_, err := o.QueryTable(&model).Filter("id", this.Id).Update(orm.Params{
		"name":    name,
		"content": content,
		"status":  status,
	})
	if err != nil {
		beego.Error(err)
		panic(vanilla.NewBusinessError("update：fail", fmt.Sprintf("修改项目失败")))
	}
}

func (this *Project) Delete() {
	_, err := vanilla.GetOrmFromContext(this.Ctx).QueryTable(m_project.Project{}).Filter(vanilla.Map{"id": this.Id}).Delete()
	if err != nil {
		beego.Error(err)
		panic(err)
	}
}

func (this *Project) AddManager(uid int) *Project {
	dbModel := &m_project.ProjectToAdministrators{
		ProjectId:       this.Id,
		AdministratorId: uid,
	}
	_, err := vanilla.GetOrmFromContext(this.Ctx).Insert(dbModel)
	if err != nil {
		beego.Error(err)
		panic(vanilla.NewSystemError("create:failed", "创建失败"))
	}
	return this
}

func (this *Project) DeleteManager(uid int) *Project {
	filter := vanilla.Map{
		"ProjectId":       this.Id,
		"AdministratorId": uid,
	}
	_, err := vanilla.GetOrmFromContext(this.Ctx).QueryTable(&m_project.ProjectToAdministrators{}).Filter(filter).Delete()
	if err != nil {
		beego.Error(err)
		panic(vanilla.NewSystemError("Delete:failed", "删除失败"))
	}
	return this
}

func (this *Project) AddMember(uid int) *Project {
	dbModel := &m_project.ProjectToParticipants{
		ProjectId:     this.Id,
		ParticipantId: uid,
	}
	_, err := vanilla.GetOrmFromContext(this.Ctx).Insert(dbModel)
	if err != nil {
		beego.Error(err)
		panic(vanilla.NewSystemError("create:failed", "创建失败"))
	}
	return this
}

func (this *Project) DeleteMember(uid int) *Project {
	filter := vanilla.Map{
		"ProjectId":     this.Id,
		"ParticipantId": uid,
	}
	_, err := vanilla.GetOrmFromContext(this.Ctx).QueryTable(&m_project.ProjectToParticipants{}).Filter(filter).Delete()
	if err != nil {
		beego.Error(err)
		panic(vanilla.NewSystemError("Delete:failed", "删除失败"))
	}
	return this
}

func (this *Project) AuthorityVerify() {
	uid := b_account.GetUserFromContext(this.Ctx).Id
	filters := vanilla.Map{
		"project_id":       this.Id,
		"administrator_id": uid,
	}
	qs := vanilla.GetOrmFromContext(this.Ctx).QueryTable(m_project.ProjectToAdministrators{})
	var model m_project.ProjectToAdministrators
	if len(filters) > 0 {
		qs = qs.Filter(filters)
	}

	err := qs.One(&model)
	if err != nil {
		beego.Error(err)
		panic(vanilla.NewBusinessError("authority verification：fail", fmt.Sprintf("非管理员不能修改项目信息")))
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
