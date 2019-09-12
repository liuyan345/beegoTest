package main

import (
	"beegoTest/models/class"
	_ "beegoTest/routers"
	"github.com/astaxie/beego"
)

func main() {
	class.TestORM()
	beego.Run()
}

