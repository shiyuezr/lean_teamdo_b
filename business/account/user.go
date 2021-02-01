package account

import (
	"context"
	"github.com/bitly/go-simplejson"
	"github.com/kfchen81/beego/vanilla"
	m_account "teamdo/models/account"
)

type User struct {
	vanilla.EntityBase
	Id             int
	Username       string
	Password       string
	EncodePassword string
	Token          string //身份识别码
	Status         int    //用户状态
	RawData        *simplejson.Json
}

func (this *User) Login() *User {
	jwtToken := vanilla.EncodeJWT(vanilla.Map{
		"type": 2,
		"uid":  this.Id,
	})
	this.Token = jwtToken
	return this
}

func NewUserFromModel(ctx context.Context, model *m_account.User) *User {
	instance := new(User)
	instance.Ctx = ctx
	instance.Id = model.Id
	instance.Username = model.Username
	instance.Password = model.Password
	return instance
}

func GetUserFromContext(ctx context.Context) *User {
	user := ctx.Value("user").(*User)
	return user
}

func init() {
}
