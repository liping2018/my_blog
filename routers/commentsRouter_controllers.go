package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["web_go/controllers:TestController"] = append(beego.GlobalControllerRouter["web_go/controllers:TestController"],
		beego.ControllerComments{
			Method: "TestRequestMessage",
			Router: `test`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
