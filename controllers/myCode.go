package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type MyCodeController struct {
	beego.Controller
}

type JSONStruct struct {
	Code int
	Msg string
}

type NestPreparer interface {
	NestPrepare()
}

func (c *MyCodeController) Prepare(){
//	if app ,ok := c.AppController.(NestPreparer);ok {
//		app.NestPrepare()
//	}
	fmt.Println("this is Prepare")
}

func (c *MyCodeController) MyString() {
	//name := c.GetString("name")
	//c.Data["name"] = name
	//c.TplName = "mycode.html"
	//fmt.Println("this is myString")
	//c.Ctx.WriteString("page code")

	//mystruct := &JSONStruct{22,"abc"}

	slice2 := make([]string, 9)
	var string string
	for i := 0; i < 9; i++ {
		string = fmt.Sprintf("%s%d",string,i)
		slice2[i] = fmt.Sprintf("%s%d","no.",i)
	}
	//mystruct :=
	c.Data["json"] = slice2
	fmt.Println(string)
	c.ServeJSON()
}