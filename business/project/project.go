package project

import (
	"context"
	b_account "teamdo/business/account"
	m_project "teamdo/models/project"
	"time"

	"github.com/kfchen81/beego/vanilla"
)

type Project struct {
	vanilla.EntityBase
	Id           int
	Name         string
	Introduction string
	Cover        string
	CreateAt     time.Time
	StartTime    time.Time
	IsEnabled    bool
	IsDeleted    bool
	//foreign key
	Users []*b_account.User

	Tasks []*Task
}

func NewProjectFromModel(ctx context.Context, model *m_project.Project) *Project {
	instance := new(Project)
	instance.Ctx = ctx
	instance.Model = model
	instance.Id = model.Id
	instance.Cover = model.Cover
	instance.CreateAt = model.CreateAt
	instance.Introduction = model.Introduction
	instance.IsDeleted = model.IsDeleted
	instance.IsEnabled = model.IsEnabled

	return instance
}

func init() {

}
