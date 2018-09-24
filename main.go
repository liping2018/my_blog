package main

import (
	_ "my_blog/routers"
	_ "my_blog/utils"

	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
