package lane

import (
	"context"
	"errors"
	"github.com/kfchen81/beego/vanilla"
	models "teamdo/models/lane"
	"github.com/kfchen81/beego"
	"github.com/kfchen81/beego/orm"
)


type  Lane struct {
	vanilla.EntityBase
	Id int
	Name string
	ProjectId int
	SortId int
	IsDelete bool
}

func NewLane(ctx context.Context,name string,projectId int,sortId int) *Lane {
	o:=vanilla.GetOrmFromContext(ctx)
	model:=models.Lane{}
	model.Name=name
	model.ProjectId=projectId
	model.SortId=sortId
	model.IsDelete=false
	id,err:= o.Insert(&model)
	if err!=nil {
		beego.Error(err)
		panic(vanilla.NewBusinessError("create_lane_failed","创建泳道失败"))
	}
	model.Id=int(id)
	return NewLaneFromDbModel(ctx,&model)
}

func (this *Lane) Delete(id int) error {
	var model models.Lane
	o := vanilla.GetOrmFromContext(this.Ctx)
	_,err:= o.QueryTable(&model).Filter("id", this.Id).Update(orm.Params{
		"is_delete":true,
	})
	if err !=nil{
		beego.Error(err)
		return errors.New("delete:delete_lane_failed")
	}
	return nil
}

func (this *Lane) Update(id int,name string,sortId int) error {
	var model models.Lane
	o := vanilla.GetOrmFromContext(this.Ctx)
	_,err:= o.QueryTable(&model).Filter("id", this.Id).Update(orm.Params{
		"name":name,
		"sort_id":sortId,
	})
	if err !=nil{
		beego.Error(err)
		return errors.New("update:update_lane_failed")
	}
	return nil
}

func  NewLaneFromDbModel(ctx context.Context, dbModel *models.Lane) *Lane {
	instance := new(Lane)
	instance.Ctx = ctx
	instance.Id = dbModel.Id
	instance.Name=dbModel.Name
	instance.ProjectId=dbModel.ProjectId
	instance.SortId = dbModel.SortId
	instance.IsDelete = dbModel.IsDelete
	return instance
}