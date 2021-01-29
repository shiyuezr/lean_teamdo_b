package account

import (
	"context"
	"time"

	"github.com/kfchen81/beego/vanilla"

	"github.com/bitly/go-simplejson"

)


type User struct {
	vanilla.EntityBase
	Id                int
	PlatformId        int
	Name              string
	Avatar            string
	Cover             string
	Sex               string
	Phone             string
	Birthday          string
	Region            string
	Slogan            string
	Longitude         float64
	Latitude          float64
	Distance          string
	Age               int64
	Code              string
	RawData           *simplejson.Json
	LastActiveTime    string
	DisplayLiveness   string
	Roles             []interface{}
	CreatedAt         time.Time
}

func NewUserFromOnlyId(ctx context.Context, id int) *User {
	user := new(User)
	user.Ctx = ctx
	user.Model = nil
	user.Id = id
	return user
}

func GetUserFromContext(ctx context.Context) *User {
	user := ctx.Value("user").(*User)
	return user
}

func (this *User) GetId() int {
	return this.Id
}

func (this *User) GetName() string {
	return this.Name
}

func (this *User) GetAvatar() string {
	return this.Avatar
}

func init() {
}
