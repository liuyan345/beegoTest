package routers

import (
	"beegoTest/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/?:id([0-9]*)", &controllers.SquareController{},"get:SquareData")
    beego.Router("/mycode", &controllers.MyCodeController{},"get:MyString")
    beego.Router("/web/login", &controllers.PublicController{},"get:Login;post:ToLogin")
    beego.Router("/web/center", &controllers.CenterController{},"get:Center;post:ToSetting")
    beego.Router("/web/logout", &controllers.CenterController{},"get:DoLogout")
    beego.Router("/web/add", &controllers.SquareController{},"get:AddPage;post:ToAdd")
    beego.Router("/web/msg/del/?:id([0-9]+)", &controllers.SquareController{},"post:Del")
    beego.Router("/web/comment/add/?:id([0-9]+)", &controllers.SquareController{},"post:CommentAdd")
    beego.Router("/web/comment/del/?:id([0-9]+)", &controllers.SquareController{},"post:CommentDel")
}
