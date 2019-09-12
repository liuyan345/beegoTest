package models

import (
	"beegoTest/models/class"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init(){
	orm.Debug = true
	orm.RegisterDriver("mysql",orm.DRMySQL)
	//orm.RegisterDataBase("default","mysql","liuyan:1234@tcp(127.0.0.1:3306)/beegoTest?charset=utf8")
	orm.RegisterDataBase("default","mysql",fmt.Sprintf("%s:%s@tcp(127.0.0.1:%s)/%s?charset=utf8",
		beego.AppConfig.String("PGY::user"),
		beego.AppConfig.String("PGY::pass"),
		beego.AppConfig.String("PGY::port"),
		beego.AppConfig.String("PGY::database"),
		))
	orm.RegisterModel(new(class.User))

	orm.RunSyncdb("default",false,true)
}
