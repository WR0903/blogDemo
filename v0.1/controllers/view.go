package controllers

import (
	"fmt"
	"myProject/models"

	"github.com/astaxie/beego"
)

type ViewController struct {
	beego.Controller
}

func (this *ViewController) Get() {
	id := string(this.Ctx.Input.Params()[":id"])
	fmt.Println("wang:", id)
	this.Data["Post"] = models.GetBlog(id)
	this.Layout = "layout.tpl"
	this.TplName = "view.tpl"
}
