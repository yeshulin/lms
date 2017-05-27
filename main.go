package main

import (
	"webproject/lms/common/function"

	_ "webproject/lms/initial"

	_ "webproject/lms/routers"

	"github.com/astaxie/beego"
)

func main() {

	beego.AddFuncMap("ConvertT", function.ConvertT)
	beego.SetStaticPath("/uploads", "uploads")
	beego.Run()
}
