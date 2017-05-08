package controllers

import (
	"github.com/astaxie/beego"
)

type WebController struct {
	beego.Controller
}

func (this *WebController) Prepare() {
	beego.LoadAppConfig("ini", "conf/web.conf")
	this.Data["Website"] = beego.AppConfig.String("Website")

}
