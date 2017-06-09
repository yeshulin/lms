package controllers

import (
	"fmt"
	"net/http"
	"webproject/4050/common/hjwt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

type WebController struct {
	beego.Controller
}

func (this *WebController) Prepare() {
	cookie := this.Ctx.GetCookie("Authorization")
	fmt.Println(cookie)
	Claims, _ := hjwt.CheckToken(cookie)
	beego.InsertFilter("/admin/*", beego.BeforeRouter, func(ctx *context.Context) {
		cookie1, err := ctx.Request.Cookie("Authorization")
		_, isok1 := hjwt.CheckToken(cookie1.Value)
		if err != nil || !isok1 {
			http.Redirect(ctx.ResponseWriter, ctx.Request, "/login", http.StatusMovedPermanently)
		}
		//		fmt.Println(cookie)
		//		fmt.Println("yeshulin")

	})
	this.Data["Claims"] = Claims
	this.Data["Website"] = beego.AppConfig.String("Website")
}
