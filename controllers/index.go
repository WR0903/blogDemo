package controllers

import (
	"myProject/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type IndexController struct {
	beego.Controller
}

func (this *IndexController) Get() {
	blogs := models.GetAll()
	this.Data["blogs"] = blogs
	logs.Info("wangran: %s", len(blogs))
	this.Layout = "layout.tpl"
	this.TplName = "index.tpl"
}
