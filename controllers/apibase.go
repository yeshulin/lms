package controllers

import (
	"github.com/astaxie/beego"
)

type ApibaseController struct {
	beego.Controller
}

func (this *ApibaseController) Prepare() {
	this.Data["Website"] = beego.AppConfig.String("Website")
}
