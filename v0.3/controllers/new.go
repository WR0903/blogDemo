package controllers

import (
	"myProject/models"
	"time"
)

type NewController struct {
	BaseController
}

func (this *NewController) Get() {
	this.TplName = "new.html"
}

func (this *NewController) Post() {
	inputs := this.Input()
	var blog models.Blog
	blog.Title = inputs.Get("title")
	blog.Tags = this.GetString("tags")
	blog.Short = this.GetString("short")
	blog.Content = inputs.Get("content")
	blog.Author = "然哥"
	blog.Created = time.Now().Format("2006-01-02 15:04:05")
	models.SaveBlog(blog)
	this.Ctx.Redirect(302, "/")
}
