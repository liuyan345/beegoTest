package controllers

import (
	"beegoTest/models/class"
	"github.com/astaxie/beego/validation"
	"strings"
)

type CenterController struct {
	BaseController
}

func (this *CenterController) Center(){
	if(this.IsLogin()){
		this.TplName = "web/center.html"
	}else{
		this.Redirect("/web/login",302)
	}
}

func (this *CenterController) ToSetting(){
	// 检测现在是否还在登录状态中
	this.CheckLogin()
	formType := this.GetString("type")
	switch formType {
	case "info":
		this.changeInfo()
	case "pwd":
		this.changePassword()
	}
}

func (this *CenterController) changeInfo(){
	// 修改用户信息
	ret := RET{
		Ok:      true,
		Content: "success",
	}

	defer func() {
		this.Data["json"] = ret
		this.ServeJSON()
	}()


	valid := validation.Validation{}
	if this.GetString("email") != "" {
		valid.Email(this.GetString("email"),"Email")
	}

	if this.GetString("phone") != "" {
		valid.Phone(this.GetString("phone"),"Phone")
	}

	valid.Required(this.GetString("nick"),"Nick")

	if this.GetString("phone") == "" && this.GetString("email") == "" {
		valid.Error("手机/邮箱不能同时为空")
	}

	switch  {
	case valid.HasErrors():
		switch valid.Errors[0].Key {
		case "Email":
			valid.Errors[0].Message = "邮箱格式错误"
		case "Nick":
			valid.Errors[0].Message = "昵称必填"
		case "Phone":
			valid.Errors[0].Message = "手机格式不正确"
		}
	default:
		user := this.GetSession("webUser").(class.User)

		// 验证新的手机、邮箱是否其他用户已存在
		if this.GetString("phone") != "" {
			checku := &class.User{}
			condition := map[string] string {"phone":strings.TrimSpace(this.GetString("phone"))}
			checku.Read(condition)
			if checku.Id != 0 && checku.Id != user.Id {
				valid.Error("手机已经被占用")
				break
			}
		}

		if this.GetString("email") != "" {
			checku := &class.User{}
			condition := map[string] string {"email":strings.TrimSpace(this.GetString("email"))}
			checku.Read(condition)
			if checku.Id != 0 && checku.Id != user.Id {
				valid.Error("邮箱已经被占用")
				break
			}
		}

		user.Nick = strings.TrimSpace(this.GetString("nick"))
		user.Phone = strings.TrimSpace(this.GetString("phone"))
		user.Email = strings.TrimSpace(this.GetString("email"))
		user.Update()
		this.DoLogin(user)// 更新session
		return
	}

	ret.Ok = false
	ret.Content = valid.Errors[0].Key + " " + valid.Errors[0].Message

	return

}

func (this *CenterController) changePassword(){
	// 修改密码信息
	ret := RET{
		Ok:      true,
		Content: "success",
	}

	defer func() {
		this.Data["json"] = ret
		this.ServeJSON()
	}()

	valid := validation.Validation{}
	valid.Required(this.GetString("old"),"Old")
	valid.Required(this.GetString("new"),"New")
	valid.Required(this.GetString("renew"),"ReNew")

	if this.GetString("renew") != this.GetString("new") {
		valid.Error("重复密码不一致")
	}
	user := this.GetSession("webUser").(class.User)

	if !PwCheck(strings.TrimSpace(this.GetString("old")),user.Password) {
		valid.Error("旧密码错误")
	}

	switch  {
	case valid.HasErrors():
		switch valid.Errors[0].Key {
		case "Old":
			valid.Errors[0].Message = "请填写旧密码"
		case "New":
			valid.Errors[0].Message = "请填写新密码"
		case "ReNew":
			valid.Errors[0].Message = "请填写重复新密码"
		}
	default:
		user.Password = PwGen(strings.TrimSpace(this.GetString("new")))
		user.Update()
		return
	}

	ret.Ok = false
	ret.Content = valid.Errors[0].Key + " " + valid.Errors[0].Message

	return
}