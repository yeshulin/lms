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

type NewsController struct {
	controllers.WebController
}

func (this *NewsController) NewsList() {
	this.TplName = "admin/news_list.html"
}

type NewsList struct {
	Id         int    `orm:"pk"`
	Catid      int    `分类ID`
	Title      string `标题`
	Catname    string `分类名称`
	Sort       int    `排序`
	Username   string `用户名`
	Addtime    int64  `添加时间`
	Updatetime int64  `更新时间`
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
	var maps []NewsList
	//fmt.Println(id)
	if id != "" {
		where = where + " and a.id = " + id
	}
	if search != "" {
		where = where + " and (a.title like '%" + search + "%')"
	}
	qb, _ := orm.NewQueryBuilder("mysql")

	// 构建查询对象
	qb.Select("a.id,a.catid,a.title,a.addtime,a.updatetime,a.sort,a.username,b.catname").
		From("news as a").
		LeftJoin("newstype as b").
		On("a.catid = b.id").
		Where(where).
		OrderBy(sort).Desc().
		Limit(ilimit).Offset(istart)

	// 导出 SQL 语句
	sql := qb.String()
	fmt.Println(sql)
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

	o := orm.NewOrm()
	news := new(models.News)
	//fmt.Println(password)
	catid, _ := strconv.Atoi(this.GetString("catid"))
	sort, _ := strconv.Atoi(this.GetString("sort"))
	status, _ := strconv.Atoi(this.GetString("status"))
	news.Catid = catid
	news.Title = this.GetString("title")
	news.Thumb = this.GetString("thumb")
	news.Keywords = this.GetString("keywords")
	news.Description = this.GetString("description")
	news.Content = this.GetString("content")
	news.Sort = sort
	news.Status = status
	news.Username = this.GetString("username")
	news.Addtime = time.Now().Unix()
	news.Updatetime = time.Now().Unix()
	id, err := o.Insert(news)
	if err != nil {
		beego.Error(err)
	}
	this.Data["json"] = map[string]interface{}{"code": "1", "message": "success!", "data": id}
	this.ServeJSON()
}

func (this *NewsController) Edit() {
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
	/*顶级分类查询*/
	qbs, _ := orm.NewQueryBuilder("mysql")
	var newstypes []models.NewsType
	qbs.Select("id,catname").
		From("newstype").
		Where("pid=0").
		OrderBy("id").Asc()
	sqls := qbs.String()
	nums, _ := o.Raw(sqls).QueryRows(&newstypes)

	this.Data["newstypes"] = newstypes
	this.Data["nums"] = nums
	this.Data["news"] = news
	this.TplName = "admin/news_edit.html"
}

func (this *NewsController) EditPost() {
	//fmt.Println(this.GetString("name"))

	o := orm.NewOrm()
	news := new(models.News)

	//fmt.Println(password)
	id, err := strconv.Atoi(this.GetString("id"))
	if err == nil {
		news.Id = id
		//fmt.Println(id)
		//fmt.Println(o.Read(&member))
		if o.Read(news) == nil {
			catid, _ := strconv.Atoi(this.GetString("catid"))
			sort, _ := strconv.Atoi(this.GetString("sort"))
			status, _ := strconv.Atoi(this.GetString("status"))
			news.Catid = catid
			news.Title = this.GetString("title")
			news.Thumb = this.GetString("thumb")
			news.Keywords = this.GetString("keywords")
			news.Description = this.GetString("description")
			news.Content = this.GetString("content")
			news.Sort = sort
			news.Status = status
			news.Username = this.GetString("username")
			news.Updatetime = time.Now().Unix()
			id, err := o.Update(news)
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
