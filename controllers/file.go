package controllers

import (
	"fmt"
	//	"strconv"
	//	"time"

	"github.com/astaxie/beego"
)

type FileController struct {
	beego.Controller
}

func (this *FileController) Upload() {
	f, h, err := this.GetFile("file")
	defer f.Close()
	//	year := strconv.Itoa(time.Now().Year())
	//	month := time.Now().Month().String()
	//	day := strconv.Itoa(time.Now().Day())
	fmt.Println(beego.AppConfig.String("UploadPath"))

	path := beego.AppConfig.String("UploadPath") + h.Filename

	if err != nil {
		beego.Error(err)
	} else {
		this.SaveToFile("file", path) // 保存位置在 static/upload, 没有文件夹要先创建
	}
	json := map[string]interface{}{"code": "1", "message": "success!", "data": path}
	this.Data["json"] = json
	this.ServeJSON()

}

func (this *FileController) UploadPage() {
	this.TplName = "uploadpage.html"
}
