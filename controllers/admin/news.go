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

type NewsController struct {
	controllers.WebController
}

func (this *NewsController) NewsList() {
	this.TplName = "admin/news_list.html"
}

func (this *NewsController) Get() {
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
	var maps []models.News
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
		From("news").
		Where(where).
		OrderBy(sort).Desc().
		Limit(ilimit).Offset(istart)

	// 导出 SQL 语句
	sql := qb.String()
	num, _ := o.Raw(sql).QueryRows(&maps)
	fmt.Println(num)
	/*查询总量*/
	qbs, _ := orm.NewQueryBuilder("mysql")
	var counts []models.News
	qbs.Select("id").
		From("news").
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
func (this *NewsController) Add() {
	/*节点查询*/
	o := orm.NewOrm()
	qbs, _ := orm.NewQueryBuilder("mysql")
	var newstype []models.NewsType
	qbs.Select("id,catname").
		From("newstype").
		Where("pid=0").
		OrderBy("id").Asc()
	sqls := qbs.String()
	nums, _ := o.Raw(sqls).QueryRows(&newstype)

	this.Data["newstype"] = newstype
	this.Data["nums"] = nums
	this.TplName = "admin/news_add.html"
}

func (this *NewsController) AddPost() {

	//	o := orm.NewOrm()
	//	news := new(models.News)
	//	//fmt.Println(password)

	//	sort, _ := strconv.Atoi(this.GetString("sort"))
	//	pid, _ := strconv.Atoi(this.GetString("pid"))
	//	news.Catname = this.GetString("catname")
	//	news.Sort = sort
	//	news.Pid = pid
	//	news.Addtime = time.Now().Unix()
	//	news.Updatetime = time.Now().Unix()
	//	id, err := o.Insert(news)
	//	if err != nil {
	//		beego.Error(err)
	//	}
	//	this.Data["json"] = map[string]interface{}{"code": "1", "message": "success!", "data": id}
	//	this.ServeJSON()
}

func (this *NewsController) Edit() {
	//	id, _ := strconv.Atoi(this.GetString("id"))
	//	o := orm.NewOrm()
	//	news := new(models.News)
	//	news.Id = id
	//	err := o.Read(news)

	//	if err == orm.ErrNoRows {
	//		fmt.Println("查询不到")
	//	} else if err == orm.ErrMissPK {
	//		fmt.Println("找不到主键")
	//	}
	//	/*顶级分类查询*/
	//	qbs, _ := orm.NewQueryBuilder("mysql")
	//	var newstypes []models.News
	//	qbs.Select("id,catname").
	//		From("news").
	//		Where("pid=0").
	//		OrderBy("id").Asc()
	//	sqls := qbs.String()
	//	nums, _ := o.Raw(sqls).QueryRows(&newstypes)

	//	this.Data["newstypes"] = newstypes
	//	this.Data["nums"] = nums
	//	this.Data["news"] = news
	//	this.TplName = "admin/news_edit.html"
}

func (this *NewsController) EditPost() {
	//	//fmt.Println(this.GetString("name"))

	//	o := orm.NewOrm()
	//	news := new(models.News)

	//	//fmt.Println(password)
	//	id, err := strconv.Atoi(this.GetString("id"))
	//	if err == nil {
	//		news.Id = id
	//		//fmt.Println(id)
	//		//fmt.Println(o.Read(&member))
	//		if o.Read(news) == nil {
	//			sort, _ := strconv.Atoi(this.GetString("sort"))
	//			pid, _ := strconv.Atoi(this.GetString("pid"))
	//			news.Catname = this.GetString("catname")
	//			news.Sort = sort
	//			news.Pid = pid
	//			news.Addtime = time.Now().Unix()
	//			news.Updatetime = time.Now().Unix()
	//			id, err := o.Update(news)
	//			if err != nil {
	//				beego.Error(err)
	//			}
	//			this.Data["json"] = map[string]interface{}{"code": "1", "message": "success!", "data": id}

	//		} else {
	//			this.Data["json"] = map[string]interface{}{"code": "0", "message": "fail!"}
	//		}
	//	}
	//	this.ServeJSON()
}

func (this *NewsController) View() {
	id, _ := strconv.Atoi(this.GetString("id"))
	o := orm.NewOrm()
	news := new(models.News)
	news.Id = id
	err := o.Read(news)

	if err == orm.ErrNoRows {
		fmt.Println("查询不到")
	} else if err == orm.ErrMissPK {
		fmt.Println("找不到主键")
	}
	this.Data["news"] = news
	this.TplName = "admin/news_view.html"
}

func (this *NewsController) Delete() {
	id, _ := strconv.Atoi(this.GetString("id"))
	o := orm.NewOrm()
	news := new(models.News)
	news.Id = id
	if num, err := o.Delete(news); err == nil {
		this.Data["json"] = map[string]interface{}{"code": "1", "message": "success!", "data": num}
	} else {
		this.Data["json"] = map[string]interface{}{"code": "0", "message": "fail!"}
	}
	this.ServeJSON()
}
