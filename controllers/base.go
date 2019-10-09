package controllers

import (
	"beegoTest/models/class"
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

func (this *BaseController) Prepare() {
	if this.IsLogin() {
		this.Data["webUser"] = this.GetSession("webUser").(class.User)
	}
}

func (this *BaseController) IsLogin() bool {
	return this.GetSession("webUser") != nil
}

func (this *BaseController) CheckLogin() {
	if !this.IsLogin() {
		this.Redirect("/web/login",302)
		this.Abort("302")
	}
}

func (this *BaseController) DoLogin(u class.User) {
	this.SetSession("webUser",u)
}

func (this *BaseController) DoLogout() {
	this.DestroySession()
	this.Redirect("/web/login",302)
}