package user

import (
	"github.com/kfchen81/beego/vanilla"
	b_user "teamdo/business/user"
)

type User struct {
	vanilla.RestResource
}

func (this *User) Resource() string {
	return "user.users"
}

func (this *User) GetParameters() map[string][]string {
	return map[string][]string{
		"GET": []string{},
	}
}

func (this *User) Get()  {
	bCtx := this.GetBusinessContext()

	filters := vanilla.Map{}
	users := b_user.NewUserRepository(bCtx).GetByFilters(filters)
	if len(users) == 0 {
		panic("用户不存在")
	}
	rUsers := b_user.NewEncodeUserService(bCtx).EncodeMany(users)
	response := vanilla.MakeResponse(vanilla.Map{"user": rUsers})
	this.ReturnJSON(response)

}
