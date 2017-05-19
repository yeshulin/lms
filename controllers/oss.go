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
	response := aliyun.Get_policy_token()
	this.Data["response"] = response
	this.TplName = "oss/webupload.html"
}
