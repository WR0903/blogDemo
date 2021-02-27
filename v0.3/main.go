package main

import (
	_ "myProject/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.BConfig.WebConfig.Session.SessionOn = true // 打开session
	beego.Run()
}
