package models

import (
	"github.com/astaxie/beego/orm"
)

type NewsType struct {
	Id         int    `orm:"pk"`
	Catname    string `分类名称`
	Pid        int    `父ID`
	Sort       int    `排序`
	Addtime    int64  `添加时间`
	Updatetime int64  `更新时间`
}

func (*NewsType) TableName() string {
	return "newstype"
}

func init() {
	orm.RegisterModel(new(NewsType))
}
