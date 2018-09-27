//file:UserController.go
//auth:lushifu
//date:2018-09-26
//desc:

package controllers

import (
	"my_blog/models"
)

type UserController struct {
	BaseController
}

func (c *UserController) Login() {

	if c.IsLogin > 0 {
		c.Redirect("/user", 302)
		return
	}

	// GET 请求
	if c.Ctx.Request.Method == "GET" {
		c.TplName = "login.html"
		return
	}

	type Post struct {
		Username, Password string
	}

	type post struct {
		Username, Password string
	}

	c.ParseForm(&post)
	ModelUser := models.AddNewUser()

}
