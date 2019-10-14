package routers

import (
	"beegoTest/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/?:id([0-9]*)", &controllers.SquareController{},"get:SquareData")//广场
    beego.Router("/mycode", &controllers.MyCodeController{},"get:MyString")
    beego.Router("/web/login", &controllers.PublicController{},"get:Login;post:ToLogin")//登录界面
    beego.Router("/web/headimg", &controllers.PublicController{},"get:HeadImg;post:ChangeHeadImg")//更新头像页面
    beego.Router("/web/center", &controllers.CenterController{},"get:Center;post:ToSetting")//个人中心设置
    beego.Router("/web/logout", &controllers.CenterController{},"get:DoLogout")// 退出登录
    beego.Router("/web/add", &controllers.SquareController{},"get:AddPage;post:ToAdd")
    beego.Router("/web/msg/del/?:id([0-9]+)", &controllers.SquareController{},"post:Del")
    beego.Router("/web/comment/add/?:id([0-9]+)", &controllers.SquareController{},"post:CommentAdd")//增加评论
    beego.Router("/web/comment/del/?:id([0-9]+)", &controllers.SquareController{},"post:CommentDel")//删除评论
}
