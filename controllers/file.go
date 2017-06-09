package controllers

import (
	//	"fmt"
	"path"
	"strconv"
	"time"

	"github.com/astaxie/beego"
)

type FileController struct {
	beego.Controller
}

func (this *FileController) Upload() {
	f, h, err := this.GetFile("file")
	if err != nil {
		json := map[string]interface{}{"code": "0", "message": "fail!"}
		this.Data["json"] = json
		beego.Error(err)
	} else {
		defer f.Close()

		//	year := strconv.Itoa(time.Now().Year())
		//	month := time.Now().Month().String()
		//	day := strconv.Itoa(time.Now().Day())
		// 获取当前年月
		filename := strconv.FormatInt(time.Now().Unix(), 10)
		//		fmt.Println(filename)
		//		fmt.Println(path.Ext(h.Filename))
		newfilename := filename + path.Ext(h.Filename)
		// 设置保存目录
		dirPath := beego.AppConfig.String("UploadPath")
		//		fmt.Println(h.Filename)

		this.SaveToFile("file", dirPath+newfilename)
		json := map[string]interface{}{"code": "1", "message": "success!", "data": "/" + dirPath + newfilename}
		this.Data["json"] = json
	}
	this.ServeJSON()

}

func (this *FileController) UploadPage() {
	this.TplName = "uploadpage.html"
}
