package user

import (
	"github.com/kfchen81/beego/vanilla"
	b_user "teamdo/business/user"
)

type User struct {
	vanilla.RestResource
}

func (this *User) Resource() string {
	return "user.user"
}

func (this *User) GetParameters() map[string][]string {
	return map[string][]string{
		"GET": []string{
			"userName",
			"passWord:int",
		},
	}
}

func (this *User) Get()  {
	bCtx := this.GetBusinessContext()
	userName := this.GetString("userName")
	passWord, _ := this.GetInt("password")

	filters := vanilla.Map{
		"userName": userName,
		"passWord": passWord,
	}
	users := b_user.NewUserRepository(bCtx).GetByFilters(filters)
	if len(users) == 0 {
		panic("用户不存在")
	}
	response := vanilla.MakeResponse(vanilla.Map{"user": users})
	this.ReturnJSON(response)

}
