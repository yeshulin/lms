package admin

import (
	"fmt"
	"strconv"
	//"time"
	"webproject/lms/controllers"
	"webproject/lms/models"

	//"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type NodeController struct {
	controllers.WebController
}

func (this *NodeController) NodeList() {
	this.TplName = "admin/node_list.html"
}

func (this *NodeController) Get() {
	//	id := this.GetString("id")
	//	limit := "10"
	//	start := this.GetString("start")
	//	page := this.GetString("page")
	//	sort := this.GetString("sortColumn")
	//	search := this.GetString("search")
	//	ilimit, _ := strconv.Atoi(limit)
	//	istart, _ := strconv.Atoi(start)
	//	where := "1=1"
	//	o := orm.NewOrm()
	//	var maps []models.Node
	//	fmt.Println(id)
	//	if id != "" {
	//		where = where + " and id = " + id
	//	}
	//	if search != "" {
	//		where = where + " and (name like '%" + search + "%')"
	//	}
	//	qb, _ := orm.NewQueryBuilder("mysql")

	//	// 构建查询对象
	//	qb.Select("id,name,status,remark,addtime,updatetime").
	//		From("node").
	//		Where(where).
	//		OrderBy(sort).Desc().
	//		Limit(ilimit).Offset(istart)

	//	// 导出 SQL 语句
	//	sql := qb.String()
	//	num, _ := o.Raw(sql).QueryRows(&maps)
	//	fmt.Println(num)
	//	/*查询总量*/
	//	qbs, _ := orm.NewQueryBuilder("mysql")
	//	var counts []models.Node
	//	qbs.Select("id,name,status,remark,addtime,updatetime").
	//		From("node").
	//		Where(where).
	//		OrderBy(sort).Desc()
	//	sqls := qbs.String()
	//	nums, _ := o.Raw(sqls).QueryRows(&counts)
	//	fmt.Println(nums)
	//	//fmt.Println(counts)
	//	//fmt.Println(maps)
	//	data := map[string]interface{}{"data": maps, "limit": limit, "page": page, "total": nums}
	//	//data["data"] = maps
	//	//json.data = maps
	//	//json.limit = limit
	//	//json.page = page
	//	//json.total = nums

	//	fmt.Println(limit)
	//	fmt.Println(page)
	//	fmt.Println(data)
	//	this.Data["json"] = data
	//	this.ServeJSON()

}
func (this *NodeController) Add() {
	/*节点查询*/
	o := orm.NewOrm()
	qbs, _ := orm.NewQueryBuilder("mysql")
	var nodes []models.Node
	qbs.Select("id,title").
		From("role").
		Where("pid=0").
		OrderBy("id").Asc()
	sqls := qbs.String()
	nums, _ := o.Raw(sqls).QueryRows(&nodes)
	this.Data["nodes"] = nodes
	this.TplName = "admin/node_add.html"
}

func (this *NodeController) AddPost() {
	//	fmt.Println(this.GetString("name"))

	//	o := orm.NewOrm()
	//	node := new(models.Node)
	//	//fmt.Println(password)
	//	node.Name = this.GetString("name")
	//	node.Status = this.GetString("status")
	//	node.Remark = this.GetString("remark")
	//	node.Addtime = time.Now().Unix()
	//	node.Updatetime = time.Now().Unix()
	//	id, err := o.Insert(node)
	//	if err != nil {
	//		beego.Error(err)
	//	}
	//	this.Data["json"] = map[string]interface{}{"code": "1", "message": "success!", "data": id}
	//	this.ServeJSON()
}

func (this *NodeController) Edit() {
	id, _ := strconv.Atoi(this.GetString("id"))
	o := orm.NewOrm()
	node := new(models.Node)
	node.Id = id
	err := o.Read(node)

	if err == orm.ErrNoRows {
		fmt.Println("查询不到")
	} else if err == orm.ErrMissPK {
		fmt.Println("找不到主键")
	}
	this.Data["node"] = node
	this.TplName = "admin/node_edit.html"
}

func (this *NodeController) EditPost() {
	//		fmt.Println(this.GetString("name"))

	//		o := orm.NewOrm()
	//		node := new(models.Node)

	//		//fmt.Println(password)
	//		id, err := strconv.Atoi(this.GetString("id"))
	//		fmt.Println(id)

	//		if err == nil {
	//			node.Id = id
	//			fmt.Println(id)
	//			//fmt.Println(o.Read(&member))
	//			if o.Read(node) == nil {
	//				node.Name = this.GetString("name")
	//				node.Status = this.GetString("status")
	//				node.Remark = this.GetString("remark")
	//				node.Addtime = time.Now().Unix()
	//				node.Updatetime = time.Now().Unix()
	//				id, err := o.Update(node)
	//				if err != nil {
	//					beego.Error(err)
	//				}
	//				this.Data["json"] = map[string]interface{}{"code": "1", "message": "success!", "data": id}

	//			} else {
	//				this.Data["json"] = map[string]interface{}{"code": "0", "message": "fail!"}
	//			}
	//		}
	//		this.ServeJSON()
}

func (this *NodeController) View() {
	id, _ := strconv.Atoi(this.GetString("id"))
	o := orm.NewOrm()
	node := new(models.Node)
	node.Id = id
	err := o.Read(node)

	if err == orm.ErrNoRows {
		fmt.Println("查询不到")
	} else if err == orm.ErrMissPK {
		fmt.Println("找不到主键")
	}
	this.Data["node"] = node
	this.TplName = "admin/node_view.html"
}

func (this *NodeController) Delete() {
	//	id, _ := strconv.Atoi(this.GetString("id"))
	//	o := orm.NewOrm()
	//	node := new(models.Node)
	//	node.Id = id
	//	if num, err := o.Delete(node); err == nil {
	//		this.Data["json"] = map[string]interface{}{"code": "1", "message": "success!", "data": num}
	//	} else {
	//		this.Data["json"] = map[string]interface{}{"code": "0", "message": "fail!"}
	//	}
	//	this.ServeJSON()
}
