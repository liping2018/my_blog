//file:filter.go
//auth:liping
//date:2018-09-25
//desc:
package filter

import (
	"fmt"
	"my_blog/utils"
	"strings"

	"github.com/astaxie/beego"

	"github.com/astaxie/beego/context"
)

var urls = []string{}

func Ignoretoken(contexUri string) bool {
	for _, uri := range urls {
		if strings.Contains(contexUri, uri) {
			return true
		}
	}
	return false
}

var userFilter = func(ctx *context.Context) {

	if !Ignoretoken(ctx.Request.RequestURI) {

		token := ctx.Input.Query("token")
		if len(token) == 0 {
			token = ctx.Input.Header("token")
		}
		//获取token失败
		if len(token) == 0 {
			fmt.Println("获取token失败")
			r := &utils.R{1, "获取token失败", make(map[string]interface{})}
			ctx.Output.JSON(r, true, false)
		}
		//验证token
		if utils.GetRedisValue(token) == "" {
			fmt.Println("验证token失败")
			r := &utils.R{1, "验证token失败", make(map[string]interface{})}
			ctx.Output.JSON(r, true, false)
		}
	}
}

func init() {
	beego.InsertFilter("/*", beego.BeforeRouter, userFilter)
}
