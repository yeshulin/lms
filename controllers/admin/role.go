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
func (this *RoleController) User() {
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
	/*用户列表*/
	var users []models.Members
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("id,username").
		From("members").
		OrderBy("id").Desc()
	sql := qb.String()
	o.Raw(sql).QueryRows(&users)
	this.Data["users"] = users
	this.Data["role"] = role
	this.TplName = "admin/role_user.html"
}

type RoleUser struct {
	Id       int    `orm:"pk"`
	Role_id  int    `角色ID`
	Rolename string `角色名称`
	User_id  int    `用户ID`
	Username string `用户名`
	Realname string `真实姓名`
}

func (this *RoleController) UserFind() {
	id, _ := strconv.Atoi(this.GetString("id"))
	limit := "10"
	start := this.GetString("start")
	page := this.GetString("page")
	//sort := this.GetString("sortColumn")
	//search := this.GetString("search")
	ilimit, _ := strconv.Atoi(limit)
	istart, _ := strconv.Atoi(start)
	var roleuser []RoleUser
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("a.id,a.role_id,a.user_id,b.name as rolename,c.username,c.realname").
		From("role_member as a").
		LeftJoin("role as b").On("a.role_id = b.id").
		LeftJoin("members as c").On("a.user_id = c.id").
		Where("a.role_id = ?").
		OrderBy("a.id").Desc().
		Limit(ilimit).Offset(istart)

	// 导出 SQL 语句
	sql := qb.String()

	// 执行 SQL 语句
	o := orm.NewOrm()
	o.Raw(sql, id).QueryRows(&roleuser)

	/*查询总量*/
	qbs, _ := orm.NewQueryBuilder("mysql")
	var counts []RoleUser
	qbs.Select("a.id,a.role_id,a.user_id,b.name as rolename,c.username,c.realname").
		From("role_member as a").
		LeftJoin("role as b").On("a.role_id = b.id").
		LeftJoin("members as c").On("a.user_id = c.id").
		Where("a.role_id = ?").
		OrderBy("a.id").Desc()
	sqls := qbs.String()
	nums, _ := o.Raw(sqls, id).QueryRows(&counts)
	fmt.Println(nums)
	data := map[string]interface{}{"data": roleuser, "limit": limit, "page": page, "total": nums}
	//data["data"] = maps
	//json.data = maps
	//json.limit = limit
	//json.page = page
	//json.total = nums

	this.Data["json"] = data
	this.ServeJSON()
}

func (this *RoleController) DelUser() {
	id, _ := strconv.Atoi(this.GetString("id"))
	o := orm.NewOrm()
	rolemember := new(models.RoleMember)
	rolemember.Id = id
	if num, err := o.Delete(rolemember); err == nil {
		this.Data["json"] = map[string]interface{}{"code": "1", "message": "success!", "data": num}
	} else {
		this.Data["json"] = map[string]interface{}{"code": "0", "message": "fail!"}
	}
	this.ServeJSON()
}

func (this *RoleController) AddUser() {
	role_id, _ := strconv.Atoi(this.GetString("role_id"))
	user_id, _ := strconv.Atoi(this.GetString("user_id"))
	o := orm.NewOrm()
	rolemember := new(models.RoleMember)
	rolemember.Role_id = role_id
	rolemember.User_id = user_id
	if num, err := o.Insert(rolemember); err == nil {
		this.Data["json"] = map[string]interface{}{"code": "1", "message": "success!", "data": num}
	} else {
		this.Data["json"] = map[string]interface{}{"code": "0", "message": "fail!"}
	}
	this.ServeJSON()
}

func (this *RoleController) Node() {
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
	/*权限列表查询*/

	this.Data["role"] = role

	var nodes []models.Node
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("id,name").
		From("node").
		Where("status=1").
		OrderBy("id").
		Desc()
	qbs := qb.String()
	nums, _ := o.Raw(qbs).QueryRows(&nodes)

	fmt.Println(nodes)
	fmt.Println(nums)
	this.TplName = "admin/role_node.html"
}

func (this *RoleController) NodeFind() {

}

func (this *RoleController) DelNode() {

}

func (this *RoleController) AddNode() {

}
