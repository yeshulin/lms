package models

import (
	"github.com/astaxie/beego/orm"
)

type RoleMember struct {
	Id      int `orm:"pk"`
	Role_id int `角色ID`
	User_id int `用户ID`
}

func (this *RoleMember) TableName() string {
	return "role_member"
}

func init() {
	orm.RegisterModel(new(RoleMember))
}
