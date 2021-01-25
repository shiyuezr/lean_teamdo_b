package main

import (
	"github.com/kfchen81/beego"
	"github.com/kfchen81/beego/orm"
	"github.com/kfchen81/beego/vanilla"
	"github.com/kfchen81/beego/vanilla/cron"
	"os"
)

// initService 初始化服务
func initService(){
}

func main() {
	orm.Debug = true

	beegoMode := os.Getenv("BEEGO_RUNMODE")
	beego.Info("BEEGO_RUNMODE: ", beegoMode)

	beego.BConfig.EnableErrorsRender = false
	beego.BConfig.RecoverFunc = vanilla.RecoverPanic
	beego.BConfig.Log.AccessLogs = true

	initService()
	
	serviceMode := vanilla.GetServiceMode()
	if serviceMode == vanilla.SERVICE_MODE_REST {
		beego.Info("run service in REST mode ")
		cron.StartCronTasks()
		defer cron.StopCronTasks()
		
		beego.Run()
	} else if serviceMode == vanilla.SERVICE_MODE_CRON {
		beego.Info("run service in CRON mode")
		cron.StartCronTasks()
		stop := make(chan int, 0)
		<-stop
	}
}

