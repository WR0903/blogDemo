package controllers

import (
	"myProject/models"
	"time"

	"github.com/astaxie/beego"
)

type EditController struct {
	beego.Controller
}

func (this *EditController) Get() {
	id := string(this.Ctx.Input.Params()[":id"])
	this.Data["Post"] = models.GetBlog(id)
	this.Layout = "layout.tpl"
	this.TplName = "edit.tpl"
}

func (this *EditController) Post() {
	inputs := this.Input()
	var blog models.Blog
	blog.Id = inputs.Get("id")
	blog.Title = inputs.Get("title")
	blog.Content = inputs.Get("content")
	blog.Created = time.Now()
	models.SaveBlog(blog)
	this.Ctx.Redirect(302, "/")
}
