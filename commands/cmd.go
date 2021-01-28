package main

import (
	"github.com/kfchen81/beego/orm"
	_ "teamdo/models"
)

func main() {
	orm.Debug = true
	orm.RunCommand()
}
