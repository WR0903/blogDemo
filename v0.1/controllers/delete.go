package controllers

import (
	"myProject/models"

	"github.com/astaxie/beego"
)

type DeleteController struct {
	beego.Controller
}

func (this *DeleteController) Get() {
	id := string(this.Ctx.Input.Params()[":id"])
	blog := models.GetBlog(id)
	this.Data["Post"] = blog
	models.DelBlog(blog)
	this.Ctx.Redirect(302, "/")
}
