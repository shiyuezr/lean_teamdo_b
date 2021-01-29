package account

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/bitly/go-simplejson"
	"github.com/kfchen81/beego"
	"github.com/kfchen81/beego/vanilla"
	m_account "teamdo/models/account"
)

type User struct {
	vanilla.EntityBase
	Id       int
	Username string
	Password string
	Token    string //身份识别码
	Status   int    //用户状态
	RawData  *simplejson.Json
}

func (this *User) Login() *User {
	base64Password := base64.StdEncoding.EncodeToString([]byte(this.Password))
	filters := vanilla.Map{
		"username": this.Username,
		"password": base64Password,
	}

	qs := vanilla.GetOrmFromContext(this.Ctx).QueryTable(m_account.User{})

	var model m_account.User
	if len(filters) > 0 {
		qs = qs.Filter(filters)
	}

	err := qs.One(&model)
	if err != nil {
		beego.Error(err)
		panic(vanilla.NewBusinessError("user:login_fail", fmt.Sprintf("登录失败，用户名或密码错误")))
		return nil
	}

	respUser := NewUserFromModel(this.Ctx, &model)
	jwtToken := vanilla.EncodeJWT(vanilla.Map{
		"type": 2,
		"uid":  respUser.Id,
	})
	respUser.Token = jwtToken
	return respUser
}

func NewLoginUser(ctx context.Context, username string, password string) *User {
	instance := new(User)
	instance.Ctx = ctx
	instance.Username = username
	instance.Password = password
	return instance
}

func NewUserFromModel(ctx context.Context, model *m_account.User) *User {
	instance := new(User)
	instance.Ctx = ctx
	instance.Id = model.Id
	instance.Username = model.Username
	instance.Password = model.Password
	return instance
}

func init() {
}
