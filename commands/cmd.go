package main

import (
	//_ "sparrow/middleware"
	//_ "sparrow/models"
	"github.com/kfchen81/beego/orm"
	_ "teamdo/models"
)

func main() {
	orm.Debug = true

	orm.RunCommand()
}
