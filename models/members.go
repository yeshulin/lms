package models

import (
	"github.com/astaxie/beego/orm"
)

type Members struct {
	Id         int    `orm:"pk"`
	Username   string `用户名`
	Password   string `密码`
	Realname   string `真实姓名`
	Email      string `邮箱`
	Phone      string `电话`
	Addtime    int64  `添加时间`
	Updatetime int64  `更新时间`
}

func (this *Members) TableName() string {
	return "members"
}

func init() {
	//orm.RegisterModel(new(Users), new(UsersProfile))
	orm.RegisterModel(new(Members))

}
