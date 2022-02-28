package models

import (
	"github.com/kfchen81/beego"
	"github.com/kfchen81/beego/orm"

	"fmt"
	_ "teamdo/models/user"
	_ "teamdo/models/project_member"
	_ "teamdo/models/project"
	_ "teamdo/models/lane"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	// set default database
	maxIdle := 50
	maxIdleLifeTime := beego.AppConfig.DefaultInt("db::DB_WAIT_TIMEOUT", 0)
	maxConn := 100

	host := beego.AppConfig.String("db::DB_HOST")
	port := beego.AppConfig.String("db::DB_PORT")
	db := beego.AppConfig.String("db::DB_NAME")
	user := beego.AppConfig.String("db::DB_USER")
	password := beego.AppConfig.String("db::DB_PASSWORD")
	charset := beego.AppConfig.String("db::DB_CHARSET")
	mysqlURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&loc=Asia%%2FShanghai", user, password, host, port, db, charset)
	
	beego.Notice("connect mysql: ", mysqlURL)
	err := orm.RegisterDataBase("default", "mysql", mysqlURL, maxIdle, maxConn, maxIdleLifeTime)
	if err != nil{
		beego.Error(err)
	}
}
