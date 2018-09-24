package controllers

import (
	"web_go/utils"

	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

func init(){


	
}

func (c *BaseController) ok() {
	r := &utils.R{0, "Success", make(map[string]interface{}, 0)}
	c.Data["json"] = r
	c.ServeJSON()
	c.StopRun()
}
