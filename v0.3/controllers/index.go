package controllers

import (
	"fmt"
	"myProject/models"
)

type IndexController struct {
	BaseController
}

func (this *IndexController) Get() {
	this.Data["blogs"] = models.GetAll()
	fmt.Println("IsLogin:", this.IsLogin, this.Loginuser)
	this.TplName = "home.html"
}
