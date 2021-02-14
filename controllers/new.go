package controllers

import (
	"myProject/models"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type NewController struct {
	beego.Controller
}

func (this *NewController) Get() {
	this.Layout = "layout.tpl"
	this.TplName = "new.tpl"
}

func (this *NewController) Post() {
	inputs := this.Input()
	var blog models.Blog
	blog.Id = models.UniqueId()
	blog.Title = inputs.Get("title")
	logs.Info("wangran: %s", inputs.Get("title"))
	blog.Content = inputs.Get("content")
	logs.Info("wangran: %s", inputs.Get("content"))
	blog.Created = time.Now()
	models.SaveBlog(blog)
	this.Ctx.Redirect(302, "/")
}
