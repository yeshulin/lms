package admin

import (
	"fmt"
	"strconv"
	"time"
	"webproject/lms/controllers"
	"webproject/lms/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type RoleController struct {
	controllers.WebController
}

func (this *RoleController) RoleList() {
	this.TplName = "admin/role_list.html"
}

func (this *RoleController) Get() {
	id := this.GetString("id")
	limit := "10"
	start := this.GetString("start")
	page := this.GetString("page")
	sort := this.GetString("sortColumn")
	search := this.GetString("search")
	ilimit, _ := strconv.Atoi(limit)
	istart, _ := strconv.Atoi(start)
	where := "1=1"
	o := orm.NewOrm()
	var maps []models.Role
	fmt.Println(id)
	if id != "" {
		where = where + " and id = " + id
	}
	if search != "" {
		where = where + " and (name like '%" + search + "%')"
	}
	qb, _ := orm.NewQueryBuilder("mysql")

	// 构建查询对象
	qb.Select("id,name,status,remark,addtime,updatetime").
		From("role").
		Where(where).
		OrderBy(sort).Desc().
		Limit(ilimit).Offset(istart)

	// 导出 SQL 语句
	sql := qb.String()
	num, _ := o.Raw(sql).QueryRows(&maps)
	fmt.Println(num)
	/*查询总量*/
	qbs, _ := orm.NewQueryBuilder("mysql")
	var counts []models.Role
	qbs.Select("id,name,status,remark,addtime,updatetime").
		From("role").
		Where(where).
		OrderBy(sort).Desc()
	sqls := qbs.String()
	nums, _ := o.Raw(sqls).QueryRows(&counts)
	fmt.Println(nums)
	//fmt.Println(counts)
	//fmt.Println(maps)
	data := map[string]interface{}{"data": maps, "limit": limit, "page": page, "total": nums}
	//data["data"] = maps
	//json.data = maps
	//json.limit = limit
	//json.page = page
	//json.total = nums

	fmt.Println(limit)
	fmt.Println(page)
	fmt.Println(data)
	this.Data["json"] = data
	this.ServeJSON()

}
func (this *RoleController) Add() {
	this.TplName = "admin/role_add.html"
}

func (this *RoleController) AddPost() {
	fmt.Println(this.GetString("name"))

	o := orm.NewOrm()
	role := new(models.Role)
	//fmt.Println(password)
	role.Name = this.GetString("name")
	role.Status = this.GetString("status")
	role.Remark = this.GetString("remark")
	role.Addtime = time.Now().Unix()
	role.Updatetime = time.Now().Unix()
	id, err := o.Insert(role)
	if err != nil {
		beego.Error(err)
	}
	this.Data["json"] = map[string]interface{}{"code": "1", "message": "success!", "data": id}
	this.ServeJSON()
}

func (this *RoleController) Edit() {
	id, _ := strconv.Atoi(this.GetString("id"))
	o := orm.NewOrm()
	role := new(models.Role)
	role.Id = id
	err := o.Read(role)

	if err == orm.ErrNoRows {
		fmt.Println("查询不到")
	} else if err == orm.ErrMissPK {
		fmt.Println("找不到主键")
	}
	this.Data["role"] = role
	this.TplName = "admin/role_edit.html"
}

func (this *RoleController) EditPost() {
	fmt.Println(this.GetString("name"))

	o := orm.NewOrm()
	role := new(models.Role)

	//fmt.Println(password)
	id, err := strconv.Atoi(this.GetString("id"))
	fmt.Println(id)

	if err == nil {
		role.Id = id
		fmt.Println(id)
		//fmt.Println(o.Read(&member))
		if o.Read(role) == nil {
			role.Name = this.GetString("name")
			role.Status = this.GetString("status")
			role.Remark = this.GetString("remark")
			role.Addtime = time.Now().Unix()
			role.Updatetime = time.Now().Unix()
			id, err := o.Update(role)
			if err != nil {
				beego.Error(err)
			}
			this.Data["json"] = map[string]interface{}{"code": "1", "message": "success!", "data": id}

		} else {
			this.Data["json"] = map[string]interface{}{"code": "0", "message": "fail!"}
		}
	}
	this.ServeJSON()
}

func (this *RoleController) View() {
	id, _ := strconv.Atoi(this.GetString("id"))
	o := orm.NewOrm()
	role := new(models.Role)
	role.Id = id
	err := o.Read(role)

	if err == orm.ErrNoRows {
		fmt.Println("查询不到")
	} else if err == orm.ErrMissPK {
		fmt.Println("找不到主键")
	}
	this.Data["role"] = role
	this.TplName = "admin/role_view.html"
}

func (this *RoleController) Delete() {
	id, _ := strconv.Atoi(this.GetString("id"))
	o := orm.NewOrm()
	role := new(models.Role)
	role.Id = id
	if num, err := o.Delete(role); err == nil {
		this.Data["json"] = map[string]interface{}{"code": "1", "message": "success!", "data": num}
	} else {
		this.Data["json"] = map[string]interface{}{"code": "0", "message": "fail!"}
	}
	this.ServeJSON()
}
