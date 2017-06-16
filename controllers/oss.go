package controllers

import (
	//"fmt"
	//	"strconv"
	//	"time"
	//"io"
	//"net/http"
	"webproject/lms/common/aliyun"

	"github.com/astaxie/beego"
)

type OssController struct {
	beego.Controller
}

func (this *OssController) WebUpload() {
	this.TplName = "oss/webupload.html"
}

func (this *OssController) Get() {
	response := aliyun.Get_policy_token()
	this.Ctx.WriteString(response)
	//	this.Data["json"] = response
	//	this.ServeJSON()
}
