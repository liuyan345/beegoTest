package routers

import (
	"beegoTest/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/mycode", &controllers.MyCodeController{},"get:MyString")
    beego.Router("/web/login", &controllers.PublicController{},"get:Login")
}
