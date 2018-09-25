package controllers

import (
	"my_blog/models"
	"my_blog/utils"

	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

func init() {

}

func (c *BaseController) ok() {
	models.AddNewUser("liping", "12345", "你好啊", "www.baidu.com", 1)
	utils.SetRedisValue("liping", "110", "50")
	r := &utils.R{0, "Success", make(map[string]interface{}, 0)}
	c.Data["json"] = r
	c.ServeJSON()
	c.StopRun()
}
