package controllers

import (
	"myProject/models"
	"time"

	"github.com/astaxie/beego"
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
	blog.Title = inputs.Get("title")
	blog.Content = inputs.Get("content")
	blog.Created = time.Now().Format("2006-01-02 15:04:05")
	models.SaveBlog(blog)
	this.Ctx.Redirect(302, "/")
}
