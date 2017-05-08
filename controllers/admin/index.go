package admin

import (
	"webproject/lms/controllers"

	//"github.com/astaxie/beego"
	//"github.com/astaxie/beego/orm"
)

type IndexController struct {
	controllers.WebController
}

func (this *IndexController) Get() {
	this.TplName = "admin/index.html"
}
func (this *IndexController) Welcome() {
	this.TplName = "admin/welcome.html"
}
