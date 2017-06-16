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

type VideosTypeController struct {
	controllers.WebController
}

func (this *VideosTypeController) NewsTypeList() {
	this.TplName = "admin/videostype_list.html"
}

func (this *VideosTypeController) Get() {
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
	var maps []models.VideosType
	//fmt.Println(id)
	if id != "" {
		where = where + " and id = " + id
	}
	if search != "" {
		where = where + " and (catname like '%" + search + "%')"
	}
	qb, _ := orm.NewQueryBuilder("mysql")

	// 构建查询对象
	qb.Select("id,catname,pid,addtime,updatetime,sort").
		From("videostype").
		Where(where).
		OrderBy(sort).Desc().
		Limit(ilimit).Offset(istart)

	// 导出 SQL 语句
	sql := qb.String()
	num, _ := o.Raw(sql).QueryRows(&maps)
	fmt.Println(num)
	/*查询总量*/
	qbs, _ := orm.NewQueryBuilder("mysql")
	var counts []models.VideosType
	qbs.Select("id").
		From("videostype").
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
func (this *VideosTypeController) Add() {
	/*节点查询*/
	o := orm.NewOrm()
	qbs, _ := orm.NewQueryBuilder("mysql")
	var videostype []models.VideosType
	qbs.Select("id,catname").
		From("videostype").
		Where("pid=0").
		OrderBy("id").Asc()
	sqls := qbs.String()
	nums, _ := o.Raw(sqls).QueryRows(&videostype)

	this.Data["videostype"] = videostype

	this.Data["nums"] = nums
	this.TplName = "admin/videostype_add.html"
}

func (this *VideosTypeController) AddPost() {

	o := orm.NewOrm()
	videostype := new(models.VideosType)
	//fmt.Println(password)

	sort, _ := strconv.Atoi(this.GetString("sort"))
	pid, _ := strconv.Atoi(this.GetString("pid"))
	videostype.Catname = this.GetString("catname")
	videostype.Sort = sort
	videostype.Pid = pid
	videostype.Addtime = time.Now().Unix()
	videostype.Updatetime = time.Now().Unix()
	id, err := o.Insert(videostype)
	if err != nil {
		beego.Error(err)
	}
	this.Data["json"] = map[string]interface{}{"code": "1", "message": "success!", "data": id}
	this.ServeJSON()
}

func (this *VideosTypeController) Edit() {
	id, _ := strconv.Atoi(this.GetString("id"))
	o := orm.NewOrm()
	videostype := new(models.VideosType)
	videostype.Id = id
	err := o.Read(videostype)

	if err == orm.ErrNoRows {
		fmt.Println("查询不到")
	} else if err == orm.ErrMissPK {
		fmt.Println("找不到主键")
	}
	/*顶级分类查询*/
	qbs, _ := orm.NewQueryBuilder("mysql")
	var videostypes []models.VideosType
	qbs.Select("id,catname").
		From("videostype").
		Where("pid=0").
		OrderBy("id").Asc()
	sqls := qbs.String()
	nums, _ := o.Raw(sqls).QueryRows(&videostypes)

	this.Data["videostypes"] = videostypes
	this.Data["nums"] = nums
	this.Data["videostype"] = videostype
	this.TplName = "admin/videostype_edit.html"
}

func (this *VideosTypeController) EditPost() {
	//fmt.Println(this.GetString("name"))

	o := orm.NewOrm()
	videostype := new(models.VideosType)

	//fmt.Println(password)
	id, err := strconv.Atoi(this.GetString("id"))
	if err == nil {
		videostype.Id = id
		//fmt.Println(id)
		//fmt.Println(o.Read(&member))
		if o.Read(videostype) == nil {
			sort, _ := strconv.Atoi(this.GetString("sort"))
			pid, _ := strconv.Atoi(this.GetString("pid"))
			videostype.Catname = this.GetString("catname")
			videostype.Sort = sort
			videostype.Pid = pid
			videostype.Updatetime = time.Now().Unix()
			id, err := o.Update(videostype)
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

func (this *VideosTypeController) View() {
	id, _ := strconv.Atoi(this.GetString("id"))
	o := orm.NewOrm()
	videostype := new(models.VideosType)
	videostype.Id = id
	err := o.Read(videostype)

	if err == orm.ErrNoRows {
		fmt.Println("查询不到")
	} else if err == orm.ErrMissPK {
		fmt.Println("找不到主键")
	}
	this.Data["videostype"] = videostype
	this.TplName = "admin/videostype_view.html"
}

func (this *VideosTypeController) Delete() {
	id, _ := strconv.Atoi(this.GetString("id"))
	o := orm.NewOrm()
	videostype := new(models.VideosType)
	videostype.Id = id
	if num, err := o.Delete(videostype); err == nil {
		this.Data["json"] = map[string]interface{}{"code": "1", "message": "success!", "data": num}
	} else {
		this.Data["json"] = map[string]interface{}{"code": "0", "message": "fail!"}
	}
	this.ServeJSON()
}
