package models

import (
	"github.com/astaxie/beego/orm"
)

type Node struct {
	Id       int    `orm:"pk"`
	Name     string `节点名称`
	Title    string `节点标题`
	Status   int    `节点状态`
	Remark   string `备注`
	Sort     int    `排序`
	Pid      int    `父节点ID`
	Level    int    `级别`
	Type     string `类型`
	Group_id int    `分组ID`
}

func (this *Node) Table() string {
	return "node"
}

func init() {
	orm.RegisterModel(new(Node))
}
