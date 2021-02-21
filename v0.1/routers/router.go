package routers

import (
	"myProject/controllers"

	"github.com/astaxie/beego"
)

func init() {
	//显示博客首页
	beego.Router("/", &controllers.IndexController{})
	//查看博客详细信息
	beego.Router("/view/:id", &controllers.ViewController{})
	//新建博客博文
	beego.Router("/new", &controllers.NewController{})
	//删除博文
	beego.Router("/delete/:id", &controllers.DeleteController{})
	//编辑博文
	beego.Router("/edit/:id", &controllers.EditController{})
}
