package models

import (
	"github.com/astaxie/beego/orm"
)

type News struct {
	Id          int    `orm:"pk"`
	Catid       int    `分类名称`
	Title       string `父ID`
	Thumb       string `排序`
	Keywords    string `关键词`
	Description string `描述`
	Content     string `内容`
	Sort        int    `排序`
	Status      int    `状态`
	Username    string `用户名`
	Addtime     int64  `添加时间`
	Updatetime  int64  `更新时间`
}

func (*News) TableName() string {
	return "news"
}

func init() {
	orm.RegisterModel(new(News))
}
