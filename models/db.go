package models

import (
	"beegoTest/models/class"
	"fmt"
	_ "github.com/Go-SQL-Driver/MySQL"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init(){
	orm.Debug = true // 打开调试模式
	//orm.DefaultTimeLoc = time.Local //设置时区
	orm.RegisterDriver("mysql", orm.DRMySQL) //设置数据库驱动方式

	// 链接数据库
	address := beego.AppConfig.String("PGY::address")
	userName := beego.AppConfig.String("PGY::user")
	datatabse := beego.AppConfig.String("PGY::database")
	password := beego.AppConfig.String("PGY::pass")
	port := beego.AppConfig.String("PGY::port")
	sqlString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&loc=%s",userName,password,address,port,datatabse,`Asia%2FShanghai`)
	orm.RegisterDataBase("default", "mysql", sqlString)

	// 注册数据表
	orm.RegisterModelWithPrefix(beego.AppConfig.String("PGY::prefix"),new(class.User),new(class.Msg),new(class.Comment)) // 注册用户数据表

	orm.RunSyncdb("default",false,true)//这里是同步数据库的问题
}
