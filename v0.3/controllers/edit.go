package controllers

import (
	"myProject/models"
	"strconv"
	"time"

	"github.com/astaxie/beego"
)

type EditController struct {
	beego.Controller
}

func (this *EditController) Get() {
	id, _ := strconv.Atoi(this.Ctx.Input.Params()[":id"])
	this.Data["Post"] = models.GetBlog(id)
	this.TplName = "edit.html"
}

func (this *EditController) Post() {
	inputs := this.Input()
	var blog models.Blog
	blog.Id, _ = strconv.Atoi(inputs.Get("id"))
	blog.Title = inputs.Get("title")
	blog.Content = inputs.Get("content")
	blog.Created = time.Now().Format("2006-01-02 15:04:05")
	models.SaveBlog(blog)
	this.Ctx.Redirect(302, "/")
}
