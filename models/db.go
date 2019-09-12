package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init(){
	orm.RegisterDriver("mysql",orm.DRMySQL)
	//orm.RegisterDataBase("bgy","mysql","liuyan:1234@tcp(127.0.0.1:3306)/beegoTest?charset=utf8mb4")
	orm.RegisterDataBase("bgy","mysql",fmt.Sprintf("%s:%s@tcp(127.0.0.1:%s)/%s?charset=utf8",
		beego.AppConfig.String("PGY::user"),
		beego.AppConfig.String("PGY::pass"),
		beego.AppConfig.String("PGY::port"),
		beego.AppConfig.String("PGY::database"),
		))
}
