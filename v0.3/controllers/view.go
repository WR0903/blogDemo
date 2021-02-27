package controllers

import (
	"myProject/models"
	"myProject/utils"
	"strconv"
)

type ViewController struct {
	BaseController
}

func (this *ViewController) Get() {
	id, _ := strconv.Atoi(this.Ctx.Input.Params()[":id"])
	blog := models.GetBlog(id)
	this.Data["Title"] = blog.Title
	this.Data["Content"] = utils.SwitchMarkdownToHtml(blog.Content)
	this.TplName = "view.html"
}
