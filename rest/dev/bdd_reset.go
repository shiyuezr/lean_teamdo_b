package dev

import (
	"github.com/kfchen81/beego"
	"github.com/kfchen81/beego/vanilla"
)

type BDDReset struct {
	vanilla.RestResource
}

func (this *BDDReset) Resource() string {
	return "dev.bdd_reset"
}

func (r *BDDReset) IsForDevTest() bool {
	return true
}

func (this *BDDReset) GetParameters() map[string][]string {
	return map[string][]string{
		"PUT": []string{},
	}
}

func (this *BDDReset) Put() {
	if beego.AppConfig.DefaultString("db::DB_HOST", "db.dev.com") != "db.dev.com" {
		panic(vanilla.NewSystemError("bdd_reset:failed", "必须连接本地数据库！(设置DB_HOST为db.dev.com，并且确保host指向本机)"))
	}
	//o := orm.NewOrm()
	//o.Raw("delete from account_account").Exec()
	response := vanilla.MakeResponse(vanilla.Map{})
	this.ReturnJSON(response)
}
