package user

import (
	"fmt"
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
		"PUT": []string{
			"user_name:string",
			"password:string",
		},
		"POST": []string{
			"user_name:string",
			"password:string",
		},
	}
}

func (this *User)Put()  {
	bCtx:=this.GetBusinessContext()
	userName:=this.GetString("user_name")
	fmt.Println(userName)
	password:=this.GetString("password")
	user:=b_user.NewUser(bCtx,userName,password)
	response := vanilla.MakeResponse(vanilla.Map{
		"id": user.Id,
	})
	this.ReturnJSON(response)
}

func (this *User)Post()  {
	bCtx:=this.GetBusinessContext()
	userName:=this.GetString("user_name")
	password:=this.GetString("password")
	user:=b_user.NewUserRepository(bCtx).GetUserByUserNameAndPassword(userName,password)
	if user==nil {
		panic(vanilla.NewSystemError("failed_login","用户名或密码错误"))
	}
	token:= vanilla.EncodeJWT(vanilla.Map{
		"type":1,
		"user_id":user.Id,
		"uid":user.Id,
	})
	response := vanilla.MakeResponse(vanilla.Map{
		"token": token,
	})
	this.ReturnJSON(response)

}