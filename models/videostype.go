package models

import (
	"github.com/astaxie/beego/orm"
)

type VideosType struct {
	Id         int    `orm:"pk"`
	Catname    string `分类名称`
	Pid        int    `父ID`
	Sort       int    `排序`
	Addtime    int64  `添加时间`
	Updatetime int64  `更新时间`
}

func (*VideosType) TableName() string {
	return "videostype"
}

func init() {
	orm.RegisterModel(new(VideosType))
}
