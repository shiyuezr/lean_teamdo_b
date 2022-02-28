package user

import (
	"context"
	"github.com/kfchen81/beego/vanilla"
	"github.com/kfchen81/beego"
	models "teamdo/models/user"
	"time"
)

type User struct {
	vanilla.EntityBase
	Id int
	UserName string
	Password string
	CreateTime time.Time
	IsDelete bool
}

func NewUser(ctx context.Context, userName string,password string) *User {
	o:=vanilla.GetOrmFromContext(ctx)
	model:=models.User{}
	model.UserName=userName
	model.Password=password
	model.IsDelete=false
	model.CreateTime=time.Now()//Format(common.TIME_LAYOUT)
	id,err:=o.Insert(&model)
	if err!=nil {
		beego.Error(err)
		panic(vanilla.NewBusinessError("create_user_failed","注册用户失败"))
	}
	model.Id=int(id)
	return NewUserFromDbModel(ctx,&model)
}

func GetUserFromContext(ctx context.Context) *User {
	user := ctx.Value("user").(*User)
	user.Ctx = ctx
	return user
}

func  NewUserFromDbModel(ctx context.Context, dbModel *models.User) *User {
	instance := new(User)
	instance.Ctx = ctx
	instance.Id = dbModel.Id
	instance.UserName=dbModel.UserName
	instance.Password = dbModel.Password
	instance.CreateTime = dbModel.CreateTime
	instance.IsDelete = dbModel.IsDelete
	return instance
}
