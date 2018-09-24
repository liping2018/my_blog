package routers

import (
	"web_go/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	//添加注解路由
	ns := beego.NewNamespace("v1",
		beego.NSNamespace("test",
			beego.NSInclude(
				&controllers.TestController{},
			),
		),
	)
	beego.AddNamespace(ns)
}



