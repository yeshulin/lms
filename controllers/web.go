package controllers

import (
	"fmt"
	"net/http"
	"webproject/lms/common/hjwt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

type WebController struct {
	beego.Controller
}

func (this *WebController) Prepare() {
	beego.InsertFilter("/admin/*", beego.BeforeRouter, func(ctx *context.Context) {
		cookie, err := ctx.Request.Cookie("Authorization")
		//fmt.Println(cookie)
		if err != nil || !hjwt.CheckToken(cookie.Value) {
			http.Redirect(ctx.ResponseWriter, ctx.Request, "/login", http.StatusMovedPermanently)
		}
	})
	beego.LoadAppConfig("ini", "conf/web.conf")
	this.Data["Website"] = beego.AppConfig.String("Website")
}
