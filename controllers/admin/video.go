package admin

import (
	"webproject/lms/controllers"
)

type VideoController struct {
	controllers.WebController
}

func (this *VideoController) VideoList() {
	this.TplName = "admin/video_list.html"
}

func (this *VideoController) Get() {

}
