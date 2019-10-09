package main

import (
	_ "beegoTest/models"
	"beegoTest/models/class"
	_ "beegoTest/routers"
	"encoding/gob"
	"github.com/astaxie/beego"
)

func init()  {
	gob.Register(class.User{})
}

func main() {
	beego.Run()
}

