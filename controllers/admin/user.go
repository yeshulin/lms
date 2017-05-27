package admin

import (
	"html/template"
	//"net/http"
	//"crypto/aes"
	"encoding/base64"
	"fmt"
	"strconv"
	"time"
	"webproject/lms/common/aesencrypt"
	"webproject/lms/common/hjwt"
	"webproject/lms/controllers"
	"webproject/lms/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	//	jwt "github.com/dgrijalva/jwt-go"
	//	"github.com/astaxie/beego/validation"
)

type UserController struct {
	controllers.WebController
}

func (this *UserController) Login() {

	this.Data["xsrfdata"] = template.HTML(this.XSRFFormHTML())
	this.TplName = "admin/login.html"
}

func (this *UserController) Post() {
	aesencrypt := new(aesencrypt.AesEncrypt)
	username := this.GetString("account")
	o := orm.NewOrm()
	member := models.Members{Username: username}
	err := o.Read(&member, "username")
	if err != nil {
		this.Data["json"] = map[string]interface{}{"code": "0", "message": "用户不存在!"}
	}
	pass, _ := aesencrypt.Encrypt(this.GetString("password"))
	//	fmt.Println(base64.StdEncoding.EncodeToString(pass))
	//	fmt.Println(member.Password)
	if base64.StdEncoding.EncodeToString(pass) != member.Password {
		this.Data["json"] = map[string]interface{}{"code": "0", "message": "用户名密码错误!"}
	} else {
		token := hjwt.GenToken(member.Id, member.Username, member.Realname, member.Email, member.Phone)
		this.Ctx.SetCookie("Authorization", token, 86400, "/")
		//		this.Ctx.Redirect(302, "/admin")
		this.Data["json"] = map[string]interface{}{"code": "1", "message": "success", "data": member.Username}
	}
	//	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	//		"sub":      member.Id,
	//		"iat":      iat,
	//		"exp":      exp,
	//		"username": member.Username,
	//		"email":    member.Email,
	//		"phone":    member.Phone,
	//		"realname": member.Realname,
	//	})

	//	tokenString, _ := token.SignedString([]byte(beego.AppConfig.String("jwtkey")))
	//	//更新登录时间，用于只允许用户一台设备登录
	//	/*this.Data["json"] = map[string]interface{}{"code": "1", "message": "success!", "tokenString": tokenString}
	//	this.ServeJSON()*/
	//	fmt.Println(tokenString)
	this.ServeJSON()

}
func (this *UserController) Reg() {
	aesencrypt := new(aesencrypt.AesEncrypt)
	password, err := aesencrypt.Encrypt(this.GetString("password"))
	if err != nil {
		beego.Error(err)
	}
	o := orm.NewOrm()
	member := new(models.Members)
	//fmt.Println(password)
	member.Username = this.GetString("username")
	member.Password = base64.StdEncoding.EncodeToString(password)
	member.Realname = this.GetString("realname")
	member.Email = this.GetString("email")
	member.Phone = this.GetString("phone")
	member.Addtime = time.Now().Unix()
	member.Updatetime = time.Now().Unix()
	id, err := o.Insert(member)
	if err != nil {
		beego.Error(err)
	}
	this.Data["json"] = map[string]interface{}{"code": "1", "message": "success!", "data": id}
	this.ServeJSON()
}

func (this *UserController) Register() {
	this.TplName = "admin/register.html"
}

func (this *UserController) UserList() {
	this.TplName = "admin/userlist.html"
}

type Counts struct {
	total int `总量`
}

func (this *UserController) Get() {
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
	var maps []models.Members
	fmt.Println(id)
	if id != "" {
		where = where + " and id = " + id
	}
	if search != "" {
		where = where + " and (username like '%" + search + "%' or realname like '%" + search + "%' or phone like '%" + search + "%' or email like '%" + search + "%')"
	}
	qb, _ := orm.NewQueryBuilder("mysql")

	// 构建查询对象
	qb.Select("id,username,realname,email,phone,addtime,updatetime").
		From("members").
		Where(where).
		OrderBy(sort).Desc().
		Limit(ilimit).Offset(istart)

	// 导出 SQL 语句
	sql := qb.String()
	num, _ := o.Raw(sql).QueryRows(&maps)
	fmt.Println(num)
	/*查询总量*/
	qbs, _ := orm.NewQueryBuilder("mysql")
	var counts []models.Members
	qbs.Select("id,username,realname,email,phone,addtime,updatetime").
		From("members").
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

func (this *UserController) Add() {
	/*用户角色*/
	o := orm.NewOrm()
	qbs, _ := orm.NewQueryBuilder("mysql")
	var roles []models.Role
	qbs.Select("id,name,status,remark,addtime,updatetime").
		From("role").
		OrderBy("id").Asc()
	sqls := qbs.String()
	nums, _ := o.Raw(sqls).QueryRows(&roles)
	fmt.Println(nums)
	this.Data["roles"] = roles
	this.TplName = "admin/user_add.html"
}

func (this *UserController) AddPost() {
	fmt.Println(this.GetString("username"))
	aesencrypt := new(aesencrypt.AesEncrypt)
	password, err := aesencrypt.Encrypt(this.GetString("password"))
	if err != nil {
		beego.Error(err)
	}
	o := orm.NewOrm()
	member := new(models.Members)
	//fmt.Println(password)
	member.Username = this.GetString("username")
	member.Password = base64.StdEncoding.EncodeToString(password)
	member.Realname = this.GetString("realname")
	member.Email = this.GetString("email")
	member.Phone = this.GetString("phone")
	member.Addtime = time.Now().Unix()
	member.Updatetime = time.Now().Unix()
	id, err := o.Insert(member)
	if err != nil {
		beego.Error(err)
	}
	role_ids, _ := strconv.Atoi(this.GetString("role_id"))
	fmt.Println(role_ids)
	res, err := o.Raw("delete from role_member where user_id = ? and role_id in(?)", id, role_ids).Exec()
	if err == nil {
		num, _ := res.RowsAffected()
		fmt.Println("mysql row affected nums: ", num)
	}
	rolemember := new(models.RoleMember)
	rolemember.User_id = int(id)
	rolemember.Role_id = role_ids
	_, err1 := o.Insert(rolemember)
	if err1 != nil {
		beego.Error(err1)
	}
	/*defer func() {
		this.Redirect("/login", 302)
	}()*/
	this.Data["json"] = map[string]interface{}{"code": "1", "message": "success!", "data": id}
	this.ServeJSON()

}

func (this *UserController) Delete() {
	id, _ := strconv.Atoi(this.GetString("id"))
	o := orm.NewOrm()
	member := new(models.Members)
	member.Id = id
	if num, err := o.Delete(member); err == nil {
		this.Data["json"] = map[string]interface{}{"code": "1", "message": "success!", "data": num}
	} else {
		this.Data["json"] = map[string]interface{}{"code": "0", "message": "fail!"}
	}
	this.ServeJSON()
}

func (this *UserController) View() {
	id, _ := strconv.Atoi(this.GetString("id"))
	o := orm.NewOrm()
	member := new(models.Members)
	member.Id = id
	err := o.Read(member)

	if err == orm.ErrNoRows {
		fmt.Println("查询不到")
	} else if err == orm.ErrMissPK {
		fmt.Println("找不到主键")
	}
	this.Data["member"] = member
	this.TplName = "admin/user_view.html"
}

func (this *UserController) Edit() {
	id, _ := strconv.Atoi(this.GetString("id"))
	o := orm.NewOrm()
	member := new(models.Members)
	member.Id = id
	err := o.Read(member)

	if err == orm.ErrNoRows {
		fmt.Println("查询不到")
	} else if err == orm.ErrMissPK {
		fmt.Println("找不到主键")
	}
	this.Data["member"] = member
	this.TplName = "admin/user_edit.html"
}

func (this *UserController) EditPost() {
	fmt.Println(this.GetString("username"))

	o := orm.NewOrm()
	member := new(models.Members)

	//fmt.Println(password)
	id, err := strconv.Atoi(this.GetString("id"))
	fmt.Println(id)

	if err == nil {
		member.Id = id
		fmt.Println(id)
		//fmt.Println(o.Read(&member))
		if o.Read(member) == nil {
			member.Username = this.GetString("username")
			password := this.GetString("password")
			if password != "" {
				aesencrypt := new(aesencrypt.AesEncrypt)
				passwords, err := aesencrypt.Encrypt(password)
				if err != nil {
					beego.Error(err)
				}
				member.Password = base64.StdEncoding.EncodeToString(passwords)
			}
			member.Realname = this.GetString("realname")
			member.Email = this.GetString("email")
			member.Phone = this.GetString("phone")
			member.Addtime = time.Now().Unix()
			member.Updatetime = time.Now().Unix()
			id, err := o.Update(member)
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
