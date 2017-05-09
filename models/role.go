package models

import (
	"github.com/astaxie/beego/orm"
)

type Role struct {
	Id         int    `orm:"pk"`
	Name       string `角色名称`
	Status     string `状态`
	Remark     string `备注`
	Addtime    int64  `添加时间`
	Updatetime int64  `更新时间`
}

func (this *Role) TableName() string {
	return "role"
}

func init() {
	//orm.RegisterModel(new(Users), new(UsersProfile))
	orm.RegisterModel(new(Role))
}
