package main

import (
	"webproject/lms/common/function"
	"webproject/lms/common/hjwt"
	_ "webproject/lms/initial"

	_ "webproject/lms/routers"

	"fmt"
	"net/http"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func main() {
	beego.InsertFilter("/admin/*", beego.BeforeRouter, func(ctx *context.Context) {
		cookie, err := ctx.Request.Cookie("Authorization")
		fmt.Println(cookie)
		if err != nil || !hjwt.CheckToken(cookie.Value) {
			http.Redirect(ctx.ResponseWriter, ctx.Request, "/login", http.StatusMovedPermanently)
		}
	})
	beego.AddFuncMap("ConvertT", function.ConvertT)
	beego.Run()
}
