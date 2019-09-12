package controllers

import "github.com/astaxie/beego"

type PublicController struct {
	beego.Controller
}

func (this *PublicController) Login(){
	this.TplName = "register/login.html"
}
